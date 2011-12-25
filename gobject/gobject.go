package gobject

/*
#ifndef uintptr
#define uintptr unsigned int*
#endif
#include <glib.h>
#include <glib-object.h>
#include <stdlib.h>
#include <stdint.h>

static inline GObject* to_GObject(void* o) { return G_OBJECT(o); }
static inline void* new_GObject(GType typ) { return g_object_new(typ, NULL); }

static inline const gchar* getTypeName(void * o) { return G_OBJECT_TYPE_NAME(o); }

static inline GObjectClass* get_object_class(void* o) {
	return G_OBJECT_GET_CLASS(o);
}

static inline GType _get_type_from_instance(void* o) {
    return G_TYPE_FROM_INSTANCE(o);
}

extern void weak_ref_callback(gpointer data, GObject* object);

static void _g_object_weak_ref(GObject* object, gint64 data) {
	gint64* intdata = (gint64*)malloc(sizeof(gint64));
	*intdata = data;
	g_object_weak_ref(object, weak_ref_callback, (gpointer)intdata);
}

static void destroy_data(gpointer data) {
	free(data);
}

static void _g_clear_object(void* object) {
	GObject* o = G_OBJECT(object);
	g_clear_object(o);
}

static GBinding* to_GBinding(void* obj) { return G_BINDING(obj); }

*/
// #cgo pkg-config: gobject-2.0
import "C"
import "unsafe"
import "reflect"
import "runtime"
import "time"
import "strings"

func GetTypeFromInstance(obj unsafe.Pointer) GType {
	return GType(C._get_type_from_instance(obj))
}

// Base interface which all gobject like structures must implement
// to be compatible with all gobject functions - public interfaces of gobject
type ObjectLike interface {
	ToNative() unsafe.Pointer
	Connect(s string, f interface{}, data ...interface{}) (*ClosureElement, *SignalError)
	Set(map[string]interface{})
	Get([]string) map[string]interface{}
}

//-----------------------------------------------------------------------
// GObject interfaces
//-----------------------------------------------------------------------
// Convenient struct - singleton
type gobject struct {
	object unsafe.Pointer
}

func (self gobject) ToNative() unsafe.Pointer {
	return self.object
}

func (self gobject) Connect(name string, f interface{}, data ...interface{}) (*ClosureElement, *SignalError) {
	return Connect(self, name, f, data...)
}

func (self gobject) Set(properties map[string]interface{}) {
	Set(self, properties)
}

func (self gobject) Get(properties []string) map[string]interface{} {
	return Get(self, properties)
}

func New(typ GType, properties map[string]interface{}) ObjectLike {
	obj := gobject{}
	obj.object = C.new_GObject(C.GType(typ))

	Set(obj, properties)

	return obj

}

func FindProperty(obj ObjectLike, propName string) *GParamSpec {
	cpn := GString(propName)
	defer cpn.Free()

	oc := C.get_object_class(obj.ToNative())
	pspec := C.g_object_class_find_property(oc, (*C.gchar)(cpn.GetPtr()))
	if pspec == nil {
		return nil
	}
	gpspec := GParamSpec{pspec}
	return &gpspec
}

func Set(obj ObjectLike, properties map[string]interface{}) {
	if len(properties) == 0 {
		return
	}

	for p, v := range properties {
		SetProperty(obj, p, v)
	}
}

func Get(obj ObjectLike, properties []string) map[string]interface{} {
	m := map[string]interface{}{}
	for _, p := range properties {
		gv := GetProperty(obj, p)
		if gv == nil {
			continue
		}
		defer gv.Free()
		val, err := ConvertToGo(gv.GetPtr(), gv.gtype)
		if err == nil {
			m[p] = val
		}
	}
	return m
}

func SetProperty(obj ObjectLike, propName string, value interface{}) {
	cpn := GString(propName)
	defer cpn.Free()
	v := ConvertToC(value)
	defer v.Free()

	C.g_object_set_property(C.to_GObject(obj.ToNative()), (*C.gchar)(cpn.GetPtr()),
		(*C.GValue)(v.ToCGValue()))

}

func GetProperty(obj ObjectLike, propName string) *GValue {
	pspec := FindProperty(obj, propName)
	if pspec == nil {
		return nil
	}
	gv := CreateCGValue(pspec.GetType())
	pn := GString(propName)
	defer pn.Free()
	C.g_object_get_property(C.to_GObject(obj.ToNative()),
		(*C.gchar)(pn.GetPtr()),
		(*C.GValue)(gv.ToCGValue()))
	return gv
}

func Ref(obj ObjectLike) {
	C.g_object_ref(C.gpointer(obj.ToNative()))
}

func Unref(obj ObjectLike) {
	C.g_object_unref(C.gpointer(obj.ToNative()))
}

func SignalLookup(name string, objectType GType) uint32 {
	// If exists, remove detail part of the signal
	var canonicalName string = name
	if index := strings.Index(name, ":"); index != -1 {
		canonicalName = name[:index]
	}
	n := GString(canonicalName)
	defer n.Free()
	s := C.g_signal_lookup((*C.gchar)(n.GetPtr()), C.GType(objectType))
	return uint32(s)
}

func GetUniqueID() int64 {
	return time.Now().UnixNano()
}

type SignalError struct {
	Err string
}

func (self SignalError) Error() string {
	return self.Err
}

func createClosure(f interface{}, data ...interface{}) ClosureFunc {
	cfunc := reflect.ValueOf(f)

	// cfunc is not function, then there is nothing to do
	if cfunc.Kind() != reflect.Func {
		return nil
	}

	// Test number and type of return arguments
	// If return value is bool and is the only one, then
	// return that value to handler.
	// For anything else automaticaly return true
	var isTrue bool = true
	if numReturn := cfunc.Type().NumOut(); numReturn == 1 {
		rt := cfunc.Type().Out(0)
		if rt.Kind() == reflect.Bool {
			isTrue = false
		}
	}

	// Do we have variadic function for additional arguments
	var argslice bool = false
	if cfunc.Type().IsVariadic() {
		argslice = true
	}

	return func(additionalArgs []interface{}) bool {
		// Additional data
		// Create slice of reflect.Value data
		var args = make([]reflect.Value, 0)
		if len(data) > 0 {
			for _, d := range data {
				args = append(args, reflect.ValueOf(d))
			}
		}
		if argslice == true {
			for _, i := range additionalArgs {
				args = append(args, reflect.ValueOf(i))
			}
		}
		res := cfunc.Call(args)
		if isTrue {
			return isTrue
		}
		return res[0].Bool()
	}
}

func CreateCustomClosure(f interface{}, data ...interface{}) (ClosureFunc, int64) {
	cFunc := createClosure(f, data...)

	return cFunc, GetUniqueID()
}

func Connect(obj ObjectLike, name string, f interface{}, data ...interface{}) (*ClosureElement, *SignalError) {
	s_id := SignalLookup(name, GetTypeFromInstance(obj.ToNative()))

	if s_id == 0 {
		return nil, &SignalError{"Signal not found"}
	}
	uid := GetUniqueID()
	c := createClosure(f, data...)
	cloEl := RegisterHandler(obj, name, uid, c)
	return cloEl, nil
}

func IsObjectFloating(obj ObjectLike) bool {
	b := C.g_object_is_floating(C.gpointer(obj.ToNative()))
	return GoBool(unsafe.Pointer(&b))
}

func RefSink(obj ObjectLike) {
	C.g_object_ref_sink(C.gpointer(obj.ToNative()))
}

func ClearObject(obj ObjectLike) {
	C._g_clear_object(obj.ToNative())
}

// WeakRef closure stack
type WeakClosureFunc func()

var _weakClosures map[int64]WeakClosureFunc

func WeakRef(obj ObjectLike, f WeakClosureFunc) int64 {
	uid := GetUniqueID()
	_weakClosures[uid] = f
	o := C.to_GObject(obj.ToNative())
	C._g_object_weak_ref(o, C.gint64(uid))

	return uid
}

//export weak_ref_callback
func weak_ref_callback(data unsafe.Pointer, obj unsafe.Pointer) {
	cint := int64(*((*C.gint64)(data)))

	if f, ok := _weakClosures[cint]; ok {
		// Call WeakClosureFunc to clean up something
		f()

		// Remove this closure from _weakClosures
		delete(_weakClosures, cint)
	}

	// Free data
	C.destroy_data(C.gpointer(data))
}

// BoxedLike interface for Boxed values
type BoxedLike interface {
	GetBoxType() GType
}

// GBinding type
type GBinding struct {
	object *C.GBinding
}

// Clear GBinding struct when it goes out of reach
func gbindingFinalizer(gb *GBinding) {
	runtime.SetFinalizer(gb, func(gb *GBinding) { Unref(gb) })
}

// Conversion funcs
func gbindingFromNative(obj unsafe.Pointer) interface{} {
	gb := &GBinding{}
	gb.object = C.to_GBinding(obj)

	if IsObjectFloating(gb) {
		RefSink(gb)
	} else {
		Ref(gb)
	}
	gbindingFinalizer(gb)

	return gb
}

func (self GBinding) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self GBinding) Connect(name string, f interface{}, data ...interface{}) (*ClosureElement, *SignalError) {
	return Connect(self, name, f, data...)
}

func (self GBinding) Set(properties map[string]interface{}) {
	Set(self, properties)
}

func (self GBinding) Get(properties []string) map[string]interface{} {
	return Get(self, properties)
}

func BindProperty(source ObjectLike, sourceProperty string, dest ObjectLike, destProperty string, g_BindingFlags int) *GBinding {
	sp := GString(sourceProperty)
	defer sp.Free()
	dp := GString(destProperty)
	defer dp.Free()
	o := C.g_object_bind_property(C.gpointer(source.ToNative()), (*C.gchar)(sp.GetPtr()),
		C.gpointer(dest.ToNative()), (*C.gchar)(dp.GetPtr()), C.GBindingFlags(g_BindingFlags))

	return gbindingFromNative(unsafe.Pointer(o)).(*GBinding)
}

var GBindingFlags gbindingFlags

type gbindingFlags struct {
	DEFAULT        int
	BIDIRECTIONAL  int
	SYNC_CREATE    int
	INVERT_BOOLEAN int
}

func init() {
	_weakClosures = make(map[int64]WeakClosureFunc)

	// Initialize GBindingFlags
	GBindingFlags.DEFAULT = 0
	GBindingFlags.BIDIRECTIONAL = 1 << 0
	GBindingFlags.SYNC_CREATE = 1 << 1
	GBindingFlags.INVERT_BOOLEAN = 1 << 2
}
