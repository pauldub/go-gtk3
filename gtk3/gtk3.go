package gtk3

/*
#include "go-gtk3.h"

static void _gtk_init(void* argc, void* argv) {
	gtk_init((int*)argc, (char***)argv);
}

// Utility funcs {{{
// Dummy callback function
static gboolean _gtk_func_handler(void* id, ...) {
	return TRUE;
}

// GClosure generator func
static inline GClosure* _go_new_closure(gint64 id) {
	gint64 *pgint64 = (gint64*)malloc(sizeof(gint64));
	*pgint64 = id;
	GClosure* c = g_cclosure_new_swap(G_CALLBACK(_gtk_func_handler),
									(gpointer)pgint64,
									_gtk_closure_destroy_id);

	g_closure_set_marshal(c, _gtk_marshal);

	return c;
}

static inline GArray* g_array_from_GValues(void* val, guint num_elements) {
	GValue* values = (GValue*)val;
	GArray* na = g_array_new(TRUE, TRUE, sizeof(GValue));
	guint i;
	for(i = 0; i < num_elements; i++) {
		g_array_append_val(na, *(values + i));
	}
	return na;
}

static inline GValue get_index(GArray* ar, guint i) {
	return g_array_index(ar, GValue, i);
}

static inline void free_array(GArray* ar) {
	g_array_free(ar, TRUE);
}
// End Utility funcs }}}


// to_Gtk*** Funcs {{{
static inline GObject* to_GObject(void* o) { return G_OBJECT(o); }
static inline GApplication* to_GApplication(GtkApplication* g) { return G_APPLICATION(g); }
static inline GtkClipboard* to_GtkClipboard(void* obj) { return GTK_CLIPBOARD(obj); }
static inline GtkAccelGroup* to_GtkAccelGroup(void* obj) { return GTK_ACCEL_GROUP(obj); }
static inline GtkWidget* to_GtkWidget(void* obj) { return GTK_WIDGET(obj); }
static inline GtkContainer* to_GtkContainer(void *obj) { return GTK_CONTAINER(obj); }
static inline GtkBin* to_GtkBin(void *obj) { return GTK_BIN(obj); }
static inline GtkWindow* to_GtkWindow(void* obj) { return GTK_WINDOW(obj); }
static inline GtkAssistant* to_GtkAssistant(void* obj) { return GTK_ASSISTANT(obj); }
static inline GtkBox* to_GtkBox(void* obj) { return GTK_BOX(obj); }
static inline GtkPaned* to_GtkPaned(void* obj) { return GTK_PANED(obj); }
static inline GtkButtonBox* to_GtkButtonBox(void* obj) { return GTK_BUTTON_BOX(obj); }
static inline GtkFrame* to_GtkFrame(void* obj) { return GTK_FRAME(obj); }
static inline GtkGrid* to_GtkGrid(void* obj) { return GTK_GRID(obj); }
static inline GtkLabel* to_GtkLabel(void* obj) { return GTK_LABEL(obj); }
static inline GtkProgressBar* to_GtkProgressBar(void* obj) { return GTK_PROGRESS_BAR(obj); }
static inline GtkStatusbar* to_GtkStatusbar(void* obj) { return GTK_STATUSBAR(obj); }
static inline GtkImage* to_GtkImage(void* obj) { return GTK_IMAGE(obj); }
static inline GtkSpinner* to_GtkSpinner(void* obj) { return GTK_SPINNER(obj); }
static inline GtkInfoBar* to_GtkInfoBar(void* obj) { return GTK_INFO_BAR(obj); }
static inline GtkButton* to_GtkButton(void* obj) { return GTK_BUTTON(obj); }
static inline GtkToggleButton* to_GtkToggleButton(void* obj) { return GTK_TOGGLE_BUTTON(obj); }
static inline GtkLinkButton* to_GtkLinkButton(void* obj) { return GTK_LINK_BUTTON(obj); }
static inline GtkCheckButton* to_GtkCheckButton(void* obj) { return GTK_CHECK_BUTTON(obj); }
static inline GtkRadioButton* to_GtkRadioButton(void* obj) { return GTK_RADIO_BUTTON(obj); }
static inline GtkSwitch* to_GtkSwitch(void* obj) { return GTK_SWITCH(obj); }
static inline GtkEntryBuffer* to_GtkEntryBuffer(void* obj) { return GTK_ENTRY_BUFFER(obj); }
static inline GtkEntryCompletion* to_GtkEntryCompletion(void* obj) { return GTK_ENTRY_COMPLETION(obj); }
static inline GtkEntry* to_GtkEntry(void* obj) { return GTK_ENTRY(obj); }
static inline GtkScale* to_GtkScale(void* obj) { return GTK_SCALE(obj); }
static inline GtkSpinButton* to_GtkSpinButton(void* obj) { return GTK_SPIN_BUTTON(obj); }
static inline GtkDialog* to_GtkDialog(void* obj) { return GTK_DIALOG(obj); }
static inline GtkMessageDialog* to_GtkMessageDialog(void* obj) { return GTK_MESSAGE_DIALOG(obj); }
static inline GtkInvisible* to_GtkInvisible(void* obj) { return GTK_INVISIBLE(obj); }
static inline GtkSeparator* to_GtkSeparator(void* obj) { return GTK_SEPARATOR(obj); }
static inline GtkAdjustment* to_GtkAdjustment(void* obj) { return GTK_ADJUSTMENT(obj); }
static inline GtkRange* to_GtkRange(void* obj) { return GTK_RANGE(obj); }
static inline GtkScrollbar* to_GtkScrollbar(void* obj) { return GTK_SCROLLBAR(obj); }
static inline GtkScrolledWindow* to_GtkScrolledWindow(void* obj) { return GTK_SCROLLED_WINDOW(obj); }
static inline GtkTextTag* to_GtkTextTag(void* obj) { return GTK_TEXT_TAG(obj); }
static inline GtkTextTagTable* to_GtkTextTagTable(void* obj) { return GTK_TEXT_TAG_TABLE(obj); }
static inline GtkTextMark* to_GtkTextMark(void* obj) { return GTK_TEXT_MARK(obj); }
static inline GtkTextBuffer* to_GtkTextBuffer(void* obj) { return GTK_TEXT_BUFFER(obj); }
static inline GtkTextView* to_GtkTextView(void* obj) { return GTK_TEXT_VIEW(obj); }
static inline GtkTreeModel* to_GtkTreeModel(void* obj) { return GTK_TREE_MODEL(obj); }
static inline GtkListStore* to_GtkListStore(void* obj) { return GTK_LIST_STORE(obj); }
static inline GtkTreeStore* to_GtkTreeStore(void* obj) { return GTK_TREE_STORE(obj); }
static inline GtkCellArea* to_GtkCellArea(void* obj) { return GTK_CELL_AREA(obj); }
static inline GtkCellAreaContext* to_GtkCellAreaContext(void* obj) { return GTK_CELL_AREA_CONTEXT(obj); }
static inline GtkCellView* to_GtkCellView(void* obj) { return GTK_CELL_VIEW(obj); }
static inline GtkCellLayout* to_GtkCellLayout(void* obj) { return GTK_CELL_LAYOUT(obj); }
static inline GtkCellRenderer* to_GtkCellRenderer(void* obj) { return GTK_CELL_RENDERER(obj); }
static inline GtkCellRendererText* to_GtkCellRendererText(void* obj) { return GTK_CELL_RENDERER_TEXT(obj); }
static inline GtkCellRendererProgress* to_GtkCellRendererProgress(void* obj) { return GTK_CELL_RENDERER_PROGRESS(obj); }
static inline GtkCellRendererSpinner* to_GtkCellRendererSpinner(void* obj) { return GTK_CELL_RENDERER_SPINNER(obj); }
static inline GtkCellRendererToggle* to_GtkCellRendererToggle(void* obj) { return GTK_CELL_RENDERER_TOGGLE(obj); }
static inline GtkCellRendererPixbuf* to_GtkCellRendererPixbuf(void* obj) { return GTK_CELL_RENDERER_PIXBUF(obj); }
static inline GtkCellRendererAccel* to_GtkCellRendererAccel(void* obj) { return GTK_CELL_RENDERER_ACCEL(obj); }
static inline GtkCellRendererCombo* to_GtkCellRendererCombo(void* obj) { return GTK_CELL_RENDERER_COMBO(obj); }
static inline GtkCellRendererSpin* to_GtkCellRendererSpin(void* obj) { return GTK_CELL_RENDERER_SPIN(obj); }
static inline GtkTreeViewColumn* to_GtkTreeViewColumn(void* obj) { return GTK_TREE_VIEW_COLUMN(obj); }
static inline GtkTreeView* to_GtkTreeView(void* obj) { return GTK_TREE_VIEW(obj); }
static inline GtkComboBox* to_GtkComboBox(void* obj) { return GTK_COMBO_BOX(obj); }
static inline GtkComboBoxText* to_GtkComboBoxText(void* obj) { return GTK_COMBO_BOX_TEXT(obj); }
static inline GtkMenuShell* to_GtkMenuShell(void* obj) { return GTK_MENU_SHELL(obj); }
static inline GtkMenu* to_GtkMenu(void* obj) { return GTK_MENU(obj); }
static inline GtkMenuBar* to_GtkMenuBar(void* obj) { return GTK_MENU_BAR(obj); }
static inline GtkMenuItem* to_GtkMenuItem(void* obj) { return GTK_MENU_ITEM(obj); }
static inline GtkCheckMenuItem* to_GtkCheckMenuItem(void* obj) { return GTK_CHECK_MENU_ITEM(obj); }
static inline GtkImageMenuItem* to_GtkImageMenuItem(void* obj) { return GTK_IMAGE_MENU_ITEM(obj); }
static inline GtkSeparatorMenuItem* to_GtkSeparatorMenuItem(void* obj) { return GTK_SEPARATOR_MENU_ITEM(obj); }
static inline GtkRadioMenuItem* to_GtkRadioMenuItem(void* obj) { return GTK_RADIO_MENU_ITEM(obj); }
static inline GtkTearoffMenuItem* to_GtkTearoffMenuItem(void* obj) { return GTK_TEAROFF_MENU_ITEM(obj); }
static inline GtkColorButton* to_GtkColorButton(void* obj) { return GTK_COLOR_BUTTON(obj); }
static inline GtkTreeSelection* to_GtkTreeSelection(void* obj) { return GTK_TREE_SELECTION(obj); }
static inline GtkNotebook* to_GtkNotebook(void* obj) { return GTK_NOTEBOOK(obj); }
static inline GtkOrientable* to_GtkOrientable(void* obj) { return GTK_ORIENTABLE(obj); }
static inline GtkAction* to_GtkAction(void* obj) { return GTK_ACTION(obj); }
// End }}}

// GtkAccelGroup funcs {{{

// End GtkAccelGroup funcs }}}

// GtkContainer funcs {{{
static void _gtk_container_foreach(GtkContainer* container, gint64* id) {
	gtk_container_foreach(container, _gtk_callback, (gpointer)id);
}

//End GtkContainer funcs }}}

// GtkApplication funcs {{{
static inline int run_app(GtkApplication *app) {
	int status = g_application_run(G_APPLICATION (app), 0, NULL);
	return status;
}

static inline GtkApplication* to_GtkApplication(void* o) {
    return GTK_APPLICATION(o);
}
// End }}}

// ButtonBox funcs {{{
static void _gtk_button_box_set_child_non_homogeneous(GtkButtonBox *widget, GtkWidget* child,
														gboolean non_homogeneous) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_button_box_set_child_non_homogeneous(widget, child, non_homogeneous);
#endif
}

static gboolean _gtk_button_box_get_child_non_homogeneous(GtkButtonBox *widget, GtkWidget* child) {
#if GTK_CHECK_VERSION(3,2,0)
	return gtk_button_box_get_child_non_homogeneous(widget, child);
#else
	return FALSE;
#endif
}
// End ButtonBox funcs }}}

// GtkGrid funcs {{{
static GtkWidget* _gtk_grid_get_child_at(GtkGrid* grid, gint left, gint top) {
#if GTK_CHECK_VERSION(3,2,0)
	return gtk_grid_get_child_at(grid, left, top);
#else
	return NULL;
#endif
}

static void _gtk_grid_insert_row(GtkGrid* grid, gint position) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_grid_insert_row(grid, position);
#endif
}

static void _gtk_grid_insert_column(GtkGrid* grid, gint position) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_grid_insert_column(grid, position);
#endif
}

static void _gtk_grid_insert_next_to(GtkGrid* grid, GtkWidget* sibling, GtkPositionType side) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_grid_insert_next_to(grid, sibling, side);
#endif
}
// End GtkGrid funcs }}}

// GtkEntry funcs {{{

static void _gtk_entry_set_placeholder_text(GtkEntry* entry, const gchar* text) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_entry_set_placeholder_text(entry, text);
#endif
}

static const gchar* _gtk_entry_get_placeholder_text(GtkEntry* entry) {
#if GTK_CHECK_VERSION(3,2,0)
	return gtk_entry_get_placeholder_text(entry);
#else
	return NULL;
#endif
}
// End GtkEntry funcs }}}


// GtkDialog funcs {{{
static inline GtkDialog* _dialog_new_with_buttons(const gchar* title,
												  GtkWindow* parent,
												  GtkDialogFlags flags,
												  const gchar* firstbutt,
												  gint resid) {
	GtkWidget* w = gtk_dialog_new_with_buttons(title, parent, flags, firstbutt, resid, NULL);
	return to_GtkDialog(w);
}
// End GtkDialog funcs }}}

// GtkMessageDialog funcs {{{
static inline GtkWidget* _new_message_dialog(GtkWindow* parent,
													GtkDialogFlags flags,
													GtkMessageType typ,
													GtkButtonsType buttons,
													const gchar* message) {
	return gtk_message_dialog_new(parent, flags, typ, buttons, message, NULL);
}

static inline void _gtk_message_dialog_format_secondary_text(GtkMessageDialog* md, const gchar* message) {
	gtk_message_dialog_format_secondary_text(md, message, NULL);
}
static inline void _gtk_message_dialog_format_secondary_markup(GtkMessageDialog* md, const gchar* message) {
	gtk_message_dialog_format_secondary_markup(md, message, NULL);
}
// End GtkMessageDialog funcs }}}

// GtkAdjustment funcs {{{
static gdouble _gtk_adjustment_get_minimum_increment(GtkAdjustment* adjustment) {
#if GTK_CHECK_VERSION(3,2,0)
    return gtk_adjustment_get_minimum_increment(adjustment);
#else
    return 0;
#endif
}
// End GtkAdjustment funcs }}}

// GtkTextTagTable funcs {{{

static void _gtk_text_tag_table_foreach(GtkTextTagTable* table, gint64 id) {
	gtk_text_tag_table_foreach(table, _gtk_text_tag_table_foreach_func, (gpointer)(&id));
}

// End GtkTextTagTable funcs }}}

// GtkTextIter funcs {{{

static gboolean _gtk_text_iter_forward_find_char(GtkTextIter* iter, const GtkTextIter* limit, gint64 data) {
	return gtk_text_iter_forward_find_char(iter, _gtk_text_char_predicate, (gpointer)(&data), limit);
}

static gboolean _gtk_text_iter_backward_find_char(GtkTextIter* iter, const GtkTextIter* limit, gint64 data) {
	return gtk_text_iter_backward_find_char(iter, _gtk_text_char_predicate, (gpointer)(&data), limit);
}

// End GtkTextIter funcs }}}

// GtkTreePath funcs {{{
static inline gint _gtk_tree_path_get_indice(gint* indices, int index) {
	return *(indices + index);
}
//End GtkTreePath func }}}

// GtkCellArea funcs {{{
static void _gtk_cell_area_foreach(GtkCellArea* area, gint64 id) {
	gint64* uid = &id;
	gtk_cell_area_foreach(area, _gtk_cell_callback, (gpointer)uid);
}

//End GtkCellArea funcs }}}

// GtkCellLayout funcs {{{
static void _gtk_cell_layout_set_cell_data_func(GtkCellLayout* layout, GtkCellRenderer* cell, gint64 id) {
	gint64* uid = (gint64*)malloc(sizeof(gint64));
	*uid = id;
	gtk_cell_layout_set_cell_data_func(layout, cell, _gtk_cell_layout_data_func, (gpointer)uid, _g_gtk_destroy_notify);
}

//End GtkCellLayout funcs }}}

// GtkTreeSelection funcs {{{
static void _gtk_tree_select_set_select_function(GtkTreeSelection* selection, gint64 id) {
	if (id == 0) {
		gtk_tree_selection_set_select_function(selection, NULL, NULL, NULL);
		return;
	}
	gint64* uid = (gint64*)malloc(sizeof(gint64));
	*uid = id;
    gtk_tree_selection_set_select_function(selection, _gtk_tree_selection_func, (gpointer)uid, _g_gtk_destroy_notify);
}
//End GtkTreeSelection funcs }}}

// GtkComboBox funcs {{{
static void _gtk_combo_box_set_row_separator_func(GtkComboBox* combo_box, gint64 id) {
	gint64* uid = (gint64*)malloc(sizeof(gint64));
	*uid = id;
	gtk_combo_box_set_row_separator_func(combo_box, _gtk_tree_view_row_separator_func, (gpointer)uid, _g_gtk_destroy_notify);
}

//End GtkComboBox funcs }}}

// GtkMenu funcs {{{
static void _gtk_menu_popup_for_device(GtkMenu* menu, GdkDevice* device, GtkWidget* pms, GtkWidget* pmi, guint button, guint32 ai, gint64 id) {
	if (id == 0) {
		gtk_menu_popup_for_device(menu, device, pms, pmi, NULL, NULL, NULL, button, ai);
		return;
	}

	gint64* uid = (gint64*)malloc(sizeof(gint64));
	*uid = id;
	gtk_menu_popup_for_device(menu, device, pms, pmi, _gtk_menu_position_func, (gpointer)uid, _g_gtk_destroy_notify, button, ai);
}

static void _gtk_menu_popup(GtkMenu* menu, GtkWidget* pms, GtkWidget* pmi, guint button, guint32 ai, gint64 id) {
	if (id == 0) {
		gtk_menu_popup(menu, pms, pmi, NULL, NULL, button, ai);
		return;
	}

	gtk_menu_popup(menu, pms, pmi, _gtk_menu_position_func, (gpointer)(&id), button, ai);
}

static void _gtk_menu_attach_to_widget(GtkMenu* menu, GtkWidget* attachWidget, gint64 id) {
	gint64* uid = (gint64*)malloc(sizeof(gint64));
	*uid = id;
	g_object_set_data_full(G_OBJECT(menu), "detachFunc", (gpointer)uid, _g_gtk_destroy_notify);

    gtk_menu_attach_to_widget(menu, attachWidget, _gtk_menu_detach_func);
}
//End GtkMenu funcs }}}

// Gtk*Store funcs {{{
static inline GType* gtype_array_new(int n) {
	return g_new0(GType, n);
}

static inline void gtype_array_free(GType* types) {
	g_free(types);
}

static inline void gtype_array_set_element(GType* types, int n, GType type) {
	types[n] = type;
}

static inline GValue* gvalue_array_new(int n) {
	return g_new0(GValue, n);
}

static inline void gvalue_array_free(GValue* vals) {
	g_free(vals);
}

static inline void gvalue_array_set_element(GValue* vals, int n, GValue* val) {
	vals[n] = *val;
}

static inline gint* gint_array_new(int n) {
	return g_new0(gint, n);
}

static inline void gint_array_free(gint* vals) {
	g_free(vals);
}

static inline void gint_array_set_element(gint* vals, int n, gint val) {
	vals[n] = val;
}

//End Gtk*Store func }}}

*/
// #cgo pkg-config: gtk+-3.0
import "C"
import "unsafe"
import "runtime"
import "fmt"
import "github.com/pauldub/go-gtk3/gobject"
//import "github.com/norisatir/go-gtk3/gobject"
import g "github.com/pauldub/go-gtk3/gobject"
//import g "github.com/norisatir/go-gtk3/gobject"
import "github.com/pauldub/go-gtk3/gdkpixbuf"
import "github.com/pauldub/go-gtk3/gdk3"
import "github.com/pauldub/go-gtk3/glib"
import "github.com/pauldub/go-gtk3/pango"

// General types and functions {{{

func Init() {
	C._gtk_init(nil, nil)
}

func Main() {
	C.gtk_main()
}

func MainQuit() {
	C.gtk_main_quit()
}

func GetCurrentEventDevice() *gdk3.Device {
    if d := C.gtk_get_current_event_device(); d != nil {
        if dev, err := gobject.ConvertToGo(unsafe.Pointer(&d)); err == nil {
            return dev.(*gdk3.Device)
        }
    }
    return nil
}

// Map variable for custom closures required by some gtk funcs
var _closures map[int64]gobject.ClosureFunc

// Convenient map for properties
type P map[string]interface{}

// Convenient struct and slice for buttons and id's (See Dialog)
type ButID struct {
	Text     string
	Response int
}
type B []ButID

// Convenient map for Columns and values (See ListStore, TreeStore)
type V map[int]interface{}

// Convenient struct and slice for TreeViewColumn attributes
type CollAttr struct {
	Attribute string
	Column    int
}
type A []CollAttr

// GtkStockItem {{{
type StockItem struct {
	StockID           string
	Label             string
	Modifier          int
	Keyval            uint
	TranslationDomain string
}

func NewStockItemFromNative(obj unsafe.Pointer) StockItem {
	si := StockItem{}
	csi := (*C.GtkStockItem)(obj)

	si.StockID = C.GoString((*C.char)(csi.stock_id))
	si.Label = C.GoString((*C.char)(csi.label))
	si.Modifier = int(csi.modifier)
	si.Keyval = uint(csi.keyval)
	si.TranslationDomain = C.GoString((*C.char)(csi.translation_domain))

	return si
}

func (self *StockItem) ToNative() unsafe.Pointer {
	var si C.GtkStockItem

	sId := gobject.GString(self.StockID)
	defer sId.Free()
	si.stock_id = C.g_strdup((*C.gchar)(sId.GetPtr()))

	sLabel := gobject.GString(self.Label)
	defer sLabel.Free()
	si.label = C.g_strdup((*C.gchar)(sLabel.GetPtr()))

	si.modifier = C.GdkModifierType(self.Modifier)
	si.keyval = C.guint(self.Keyval)

	sTD := gobject.GString(self.TranslationDomain)
	defer sTD.Free()
	si.translation_domain = C.g_strdup((*C.gchar)(sTD.GetPtr()))

	return unsafe.Pointer(&si)
}

func StockLookup(stockId string) (bool, StockItem) {
	var ci C.GtkStockItem

	s := gobject.GString(stockId)
	defer s.Free()

	b := C.gtk_stock_lookup((*C.gchar)(s.GetPtr()), &ci)
	gb := gobject.GoBool(unsafe.Pointer(&b))

	return gb, NewStockItemFromNative(unsafe.Pointer(&ci))
}

// End StockItem }}}

////////////////////////////// }}}

// Interfaces {{{
//////////////////////////////

// Widget-like interface must have method W()
type WidgetLike interface {
	W() *Widget
	gobject.ObjectLike
}

// Container-like interface must have method C()
type ContainerLike interface {
	C() *Container
	WidgetLike
}

// Bin-like interface must have method CBin()
type BinLike interface {
	CBin() *Bin
}

// Window-like interface must have method Wnd()
type WindowLike interface {
	Wnd() *Window
	ContainerLike
}

// Box-like interface must have method CBox()
type BoxLike interface {
	CBox() *Box
}

// Button-like interface must have method B()
type ButtonLike interface {
	B() *Button
}

// Dialog-like interface must have method D()
type DialogLike interface {
	D() *Dialog
}

// Range-like interface must have method R()
type RangeLike interface {
	R() *Range
}

// TreeModel like interface must have method ITreeModel()
type TreeModelLike interface {
	ITreeModel() *TreeModel
}

// CellRendererLike interface must have method CRenderer()
type CellRendererLike interface {
	CRenderer() *CellRenderer
}

// CellAreaLike interface must have method CArea()
type CellAreaLike interface {
	CArea() *CellArea
}

// CellLayoutLike interface must have method ICellLayout
type CellLayoutLike interface {
	ICellLayout() *CellLayout
}

// ComboLike interface must have method Combo
type ComboLike interface {
	Combo() *ComboBox
}

// MenuShellLike interface must have method MShell
type MenuShellLike interface {
	MShell() *MenuShell
}

// MenuItemLike interface must have method MItem
type MenuItemLike interface {
	MItem() *MenuItem
}

// OrientableLike interface must have method IOrientable
type OrientableLike interface {
	IOrientable() *Orientable
}

// EntryLike interface must have method E()
type EntryLike interface {
	E() *Entry
}

//////////////////////////////
// END Interfaces
////////////////////////////// }}}

// GTK3 Core {{{
//////////////////////////////

// GtkAccelGroup {{{
//////////////////////////////

// GtkAccelGroup type
type AccelGroup struct {
	object *C.GtkAccelGroup
}

func NewAccelGroup() *AccelGroup {
	ag := &AccelGroup{}
	ag.object = C.gtk_accel_group_new()

	if gobject.IsObjectFloating(ag) {
		gobject.RefSink(ag)
	} else {
		gobject.Ref(ag)
	}
	accelGroupFinalizer(ag)

	return ag
}

func AccelGroupsActivate(obj gobject.ObjectLike, accelKey uint, gdk_modifier int) bool {
	o := C.to_GObject(obj.ToNative())
	b := C.gtk_accel_groups_activate(o, C.guint(accelKey), C.GdkModifierType(gdk_modifier))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func AccelGroupsFromObject(obj gobject.ObjectLike) *glib.GSList {
	o := C.to_GObject(obj.ToNative())
	clist := C.gtk_accel_groups_from_object(o)

	if clist == nil {
		return nil
	}

	goList := glib.NewGSListFromNative(unsafe.Pointer(clist))
	goList.ConversionFunc = newAccelGroupFromNative

	return goList
}

// Clear AccelGroup struct when it goes out of reach
func accelGroupFinalizer(ag *AccelGroup) {
	runtime.SetFinalizer(ag, func(ag *AccelGroup) { gobject.Unref(ag) })
}

// Conversion funcs
func newAccelGroupFromNative(obj unsafe.Pointer) interface{} {
	ag := &AccelGroup{}
	ag.object = C.to_GtkAccelGroup(obj)

	if gobject.IsObjectFloating(ag) {
		gobject.RefSink(ag)
	} else {
		gobject.Ref(ag)
	}
	accelGroupFinalizer(ag)

	return ag
}

func nativeFromAccelGroup(ag interface{}) *gobject.GValue {
	accel, ok := ag.(*AccelGroup)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ACCEL_GROUP, accel.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self AccelGroup) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self AccelGroup) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self AccelGroup) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self AccelGroup) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// AccelGroup interface

func (self *AccelGroup) ConnectAccel(accelKey uint, gtk_Modifier, gtk_AccelFlags int, f interface{}, data ...interface{}) {
	call, id := gobject.CreateCustomClosure(f, data...)
	_closures[id] = call

	gClosure := C._go_new_closure(C.gint64(id))

	C.gtk_accel_group_connect(self.object, C.guint(accelKey), C.GdkModifierType(gtk_Modifier),
		C.GtkAccelFlags(gtk_AccelFlags), gClosure)
}

func (self *AccelGroup) ConnectByPath(accelPath string, f interface{}, data ...interface{}) {
	call, id := gobject.CreateCustomClosure(f, data...)
	_closures[id] = call

	gClosure := C._go_new_closure(C.gint64(id))

	path := gobject.GString(accelPath)
	defer path.Free()

	C.gtk_accel_group_connect_by_path(self.object, (*C.gchar)(path.GetPtr()), gClosure)
}

func (self *AccelGroup) DisconectKey(accelKey uint, gdk_modifier int) bool {
	b := C.gtk_accel_group_disconnect_key(self.object, C.guint(accelKey), C.GdkModifierType(gdk_modifier))
	return gobject.GoBool(unsafe.Pointer(&b))
}

//TODO: gtk_accel_group_query
//TODO: gtk_accel_group_activate

func (self *AccelGroup) Lock() {
	C.gtk_accel_group_lock(self.object)
}

func (self *AccelGroup) UnLock() {
	C.gtk_accel_group_unlock(self.object)
}

func (self *AccelGroup) GetIsLocked() bool {
	b := C.gtk_accel_group_get_is_locked(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *AccelGroup) GetModifierMask() int {
	return int(C.gtk_accel_group_get_modifier_mask(self.object))
}

//////////////////////////////
// End GtkAccelGroup
////////////////////////////// }}}

// GtkClipboard {{{
//////////////////////////////

// GtkClipboard type
type Clipboard struct {
	object *C.GtkClipboard
}

func ClipboardGet(atom gdk3.Atom) *Clipboard {
	if atom.ToNative() == nil {
		return nil
	}
	c := &Clipboard{}
	c.object = C.gtk_clipboard_get(*((*C.GdkAtom)(atom.ToNative())))
	gobject.Ref(c)
	clipboardFinalizer(c)

	return c
}

func ClipboardGetForDisplay(display *gdk3.Display, atom gdk3.Atom) *Clipboard {
	if display == nil || atom.ToNative() == nil {
		return nil
	}

	c := &Clipboard{}
	c.object = C.gtk_clipboard_get_for_display((*C.GdkDisplay)(display.ToNative()), *((*C.GdkAtom)(atom.ToNative())))
	gobject.Ref(c)
	clipboardFinalizer(c)

	return c
}

func newClipboardFromNative(obj unsafe.Pointer) interface{} {
	c := &Clipboard{}
	c.object = C.to_GtkClipboard(obj)

	gobject.Ref(c)
	clipboardFinalizer(c)

	return c
}

// Clear Clipboard struct when it goes out of reach
func clipboardFinalizer(c *Clipboard) {
	runtime.SetFinalizer(c, func(c *Clipboard) { gobject.Unref(c) })
}

// To be object-like
func (self Clipboard) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Clipboard) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Clipboard) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Clipboard) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Clipboard interface
//TODO: Clipboard methods
//////////////////////////////
// End GtkClipboard
////////////////////////////// }}}

//////////////////////////////
// End GTK3 Core
////////////////////////////// }}}

// Base Structs {{{

// GtkWidget {{{
//////////////////////////////

// Requisition type
type Requisition struct {
	Width, Height int
}

func (self Requisition) ToNative() unsafe.Pointer {
	var req C.GtkRequisition
	req.width = C.gint(self.Width)
	req.height = C.gint(self.Height)
	return unsafe.Pointer(&req)
}

// Widget Type
type Widget struct {
	object *C.GtkWidget
}

func NewWidget(o unsafe.Pointer) *Widget {
	w := &Widget{C.to_GtkWidget(o)}

	if gobject.IsObjectFloating(w) {
		gobject.RefSink(w)
	} else {
		gobject.Ref(w)
	}
	widgetFinalizer(w)

	return w
}

func newWidgetFromNative(obj unsafe.Pointer) interface{} {
	w := &Widget{}
	w.object = C.to_GtkWidget(obj)

	if gobject.IsObjectFloating(w) {
		gobject.RefSink(w)
	} else {
		gobject.Ref(w)
	}
	widgetFinalizer(w)

	return w
}

// Clear Widget struct when it goes out of reach
func widgetFinalizer(w *Widget) {
	runtime.SetFinalizer(w, func(w *Widget) { gobject.Unref(w) })
}

// To be object-like
func (self Widget) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Widget) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Widget) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Widget) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Widget inteface

func (self *Widget) Destroy() {
	C.gtk_widget_destroy(self.object)
	runtime.SetFinalizer(self, nil)
}

func (self *Widget) InDestruction() bool {
	b := C.gtk_widget_in_destruction(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) Show() {
	C.gtk_widget_show(self.object)
}

func (self *Widget) ShowNow() {
	C.gtk_widget_show_now(self.object)
}

func (self *Widget) Hide() {
	C.gtk_widget_hide(self.object)
}

func (self *Widget) ShowAll() {
	C.gtk_widget_show_all(self.object)
}

//TODO: gtk_widget_draw

func (self *Widget) QueueDraw() {
	C.gtk_widget_queue_draw(self.object)
}

func (self *Widget) AddAccelerator(accelSignal string, accelGroup *AccelGroup, accelKey uint,
	gdk_Modifier int, gtk_AccelFlags int) {

	ac := gobject.GString(accelSignal)
	defer ac.Free()
	C.gtk_widget_add_accelerator(self.object, (*C.gchar)(ac.GetPtr()), accelGroup.object,
		C.guint(accelKey), C.GdkModifierType(gdk_Modifier), C.GtkAccelFlags(gtk_AccelFlags))
}

func (self *Widget) RemoveAccelerator(accelGroup *AccelGroup, accelKey uint, gdk_Modifier int) bool {
	b := C.gtk_widget_remove_accelerator(self.object, accelGroup.object, C.guint(accelKey), C.GdkModifierType(gdk_Modifier))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) SetAccelPath(accelPath string, accelGroup *AccelGroup) {
	s := gobject.GString(accelPath)
	defer s.Free()
	C.gtk_widget_set_accel_path(self.object, (*C.gchar)(s.GetPtr()), accelGroup.object)
}

func (self *Widget) CanActivateAccel(signalId uint) bool {
	b := C.gtk_widget_can_activate_accel(self.object, C.guint(signalId))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) Activate() bool {
	b := C.gtk_widget_activate(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) Reparent(new_parent WidgetLike) {
	C.gtk_widget_reparent(self.object, new_parent.W().object)
}

func (self *Widget) Intersect(area gdk3.Rectangle) (bool, gdk3.Rectangle) {
	var rec C.GdkRectangle
	cArea := gobject.ConvertToC(area)
	defer cArea.Free()

	b := C.gtk_widget_intersect(self.object, (*C.GdkRectangle)(cArea.GetPtr()), &rec)
	retBool := gobject.GoBool(unsafe.Pointer(&b))

	if rectangle, err := gobject.ConvertToGo(unsafe.Pointer(&rec), gdk3.GdkType.RECTANGLE); err == nil {
		return retBool, rectangle.(gdk3.Rectangle)
	}

	return retBool, gdk3.Rectangle{}
}

func (self *Widget) IsFocus() bool {
	b := C.gtk_widget_is_focus(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(self.object)
}

func (self *Widget) GrabDefault() {
	C.gtk_widget_grab_default(self.object)
}

func (self *Widget) SetName(name string) {
	gobject.SetProperty(self, "name", name)
}

func (self *Widget) GetName() string {
	n := gobject.GetProperty(self, "name")
	defer n.Free()
	return gobject.GoString(n.GetPtr())
}

func (self *Widget) SetSensitive(sensitive bool) {
	gobject.SetProperty(self, "sensitive", sensitive)
}

func (self *Widget) SetParent(parent WidgetLike) {
	C.gtk_widget_set_parent(self.object, parent.W().object)
}

func (self *Widget) SetParentWindow(parentWindow *gdk3.Window) {
	C.gtk_widget_set_parent_window(self.object, (*C.GdkWindow)(parentWindow.ToNative()))
}

func (self *Widget) GetParentWindow() *gdk3.Window {
	pw := C.gtk_widget_get_parent_window(self.object)

	if parWindow, err := gobject.ConvertToGo(unsafe.Pointer(pw)); err == nil {
		return parWindow.(*gdk3.Window)
	}
	return nil
}

func (self *Widget) SetEvents(gdk_EventMask int) {
	C.gtk_widget_set_events(self.object, C.gint(gdk_EventMask))
}

func (self *Widget) GetEvents() int {
	return int(C.gtk_widget_get_events(self.object))
}

func (self *Widget) AddEvents(gdk_EventMask int) {
	C.gtk_widget_add_events(self.object, C.gint(gdk_EventMask))
}

func (self *Widget) SetDeviceEvents(device *gdk3.Device, gdk_EventMask int) {
	C.gtk_widget_set_device_events(self.object, (*C.GdkDevice)(device.ToNative()), C.GdkEventMask(gdk_EventMask))
}

func (self *Widget) GetDeviceEvents(device *gdk3.Device) int {
	return int(C.gtk_widget_get_device_events(self.object, (*C.GdkDevice)(device.ToNative())))
}

func (self *Widget) AddDeviceEvents(device *gdk3.Device, gdk_EventMask int) {
	C.gtk_widget_add_device_events(self.object, (*C.GdkDevice)(device.ToNative()), C.GdkEventMask(gdk_EventMask))
}

func (self *Widget) SetDeviceEnabled(device *gdk3.Device, enabled bool) {
	b := gobject.GBool(enabled)
	defer b.Free()
	C.gtk_widget_set_device_enabled(self.object, (*C.GdkDevice)(device.ToNative()), *((*C.gboolean)(b.GetPtr())))
}

func (self *Widget) GetDeviceEnabled(device *gdk3.Device) bool {
	b := C.gtk_widget_get_device_enabled(self.object, (*C.GdkDevice)(device.ToNative()))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) GetTopLevel() WidgetLike {
	w := C.gtk_widget_get_toplevel(self.object)
	if tl, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
		return tl.(WidgetLike)
	}
	return nil
}

func (self *Widget) GetAncestor(widget_type gobject.GType) WidgetLike {
	w := C.gtk_widget_get_ancestor(self.object, C.GType(widget_type))
	if w == nil {
		return nil
	}
	if par, err := gobject.ConvertToGo(unsafe.Pointer(&w)); err == nil {
		return par.(WidgetLike)
	}
	return nil
}

//TODO: gdk_widget_get_visual
//TODO: gdk_widget_set_visual

func (self *Widget) IsAncestor(ancestor WidgetLike) bool {
	b := C.gtk_widget_is_ancestor(self.object, ancestor.W().object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) TranslateCoordinates(destWidget WidgetLike, srcX, srcY int) (b bool, destX, destY int) {
	var cx, cy C.gint
	cb := C.gtk_widget_translate_coordinates(self.object, destWidget.W().object, C.gint(srcX), C.gint(srcY), &cx, &cy)
	b = gobject.GoBool(unsafe.Pointer(&cb))
	destX = int(cx)
	destY = int(cy)
	return
}

func (self *Widget) HideOnDelete() bool {
	b := C.gtk_widget_hide_on_delete(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) SetDirection(gtk_TextDirection int) {
	C.gtk_widget_set_direction(self.object, C.GtkTextDirection(gtk_TextDirection))
}

func (self *Widget) GetDirection() int {
	return int(C.gtk_widget_get_direction(self.object))
}

func (self *Widget) SetDefaultDirection(gtk_TextDirection int) {
	C.gtk_widget_set_default_direction(C.GtkTextDirection(gtk_TextDirection))
}

func (self *Widget) GetDefaultDirection() int {
	return int(C.gtk_widget_get_default_direction())
}

func (self *Widget) GetCompositeName() string {
	s := C.gtk_widget_get_composite_name(self.object)
	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}

func (self *Widget) OverrideBackgroundColor(gtk_StateFlags int, color *gdk3.RGBA) {
	if color == nil {
		C.gtk_widget_override_background_color(self.object, C.GtkStateFlags(gtk_StateFlags), nil)
		return
	}
	cl := gobject.ConvertToC(*color)
	defer cl.Free()
	C.gtk_widget_override_background_color(self.object, C.GtkStateFlags(gtk_StateFlags), (*C.GdkRGBA)(cl.GetPtr()))
}

func (self *Widget) OverrideColor(gtk_StateFlags int, color *gdk3.RGBA) {
	if color == nil {
		C.gtk_widget_override_color(self.object, C.GtkStateFlags(gtk_StateFlags), nil)
		return
	}
	cl := gobject.ConvertToC(*color)
	defer cl.Free()
	C.gtk_widget_override_color(self.object, C.GtkStateFlags(gtk_StateFlags), (*C.GdkRGBA)(cl.GetPtr()))
}

func (self *Widget) OverrideFont(fontDesc *pango.FontDescription) {
	if fontDesc == nil {
		C.gtk_widget_override_font(self.object, nil)
		return
	}
	C.gtk_widget_override_font(self.object, (*C.PangoFontDescription)(fontDesc.ToNative()))
}

func (self *Widget) OverrideSymbolicColor(name string, color *gdk3.RGBA) {
	s := gobject.GString(name)
	defer s.Free()
	if color == nil {
		C.gtk_widget_override_symbolic_color(self.object, (*C.gchar)(s.GetPtr()), nil)
		return
	}
	cl := gobject.ConvertToC(*color)
	defer cl.Free()
	C.gtk_widget_override_symbolic_color(self.object, (*C.gchar)(s.GetPtr()), (*C.GdkRGBA)(cl.GetPtr()))
}

func (self *Widget) OverrideCursor(cursor, secondaryCursor *gdk3.RGBA) {
	var c, sc *C.GdkRGBA
	if cursor != nil {
		cur := gobject.ConvertToC(*cursor)
		defer cur.Free()
		c = (*C.GdkRGBA)(cur.GetPtr())
	}

	if secondaryCursor != nil {
		cur1 := gobject.ConvertToC(*secondaryCursor)
		defer cur1.Free()
		sc = (*C.GdkRGBA)(cur1.GetPtr())
	}
	C.gtk_widget_override_cursor(self.object, c, sc)
}

func (self *Widget) RenderIconPixbuf(stockId string, gtk_IconSize int) *gdkpixbuf.Pixbuf {
	s := gobject.GString(stockId)
	defer s.Free()
	p := C.gtk_widget_render_icon_pixbuf(self.object, (*C.gchar)(s.GetPtr()), C.GtkIconSize(gtk_IconSize))

	if p != nil {
		if pix, err := gobject.ConvertToGo(unsafe.Pointer(p)); err == nil {
			return pix.(*gdkpixbuf.Pixbuf)
		}
	}
	return nil
}

func (self *Widget) PopCompositeChild() {
	C.gtk_widget_pop_composite_child()
}

func (self *Widget) PushCompositeChild() {
	C.gtk_widget_push_composite_child()
}

func (self *Widget) GetParent() WidgetLike {
	w := C.gtk_widget_get_parent(self.object)
	if w != nil {
		if wid, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
			return wid.(WidgetLike)
		}
	}
	return nil
}

func (self *Widget) GetScreen() *gdk3.Screen {
	s := C.gtk_widget_get_screen(self.object)

	if s != nil {
		if scr, err := gobject.ConvertToGo(unsafe.Pointer(s)); err == nil {
			return scr.(*gdk3.Screen)
		}
	}
	return nil
}

func (self *Widget) GetWindow() *gdk3.Window {
    w := C.gtk_widget_get_window(self.object)
    if w != nil {
        if win, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
            return win.(*gdk3.Window)
        }
    }
    return nil
}

func (self *Widget) GetVisible() bool {
	b := C.gtk_widget_get_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) SetVisible(visible bool) {
	b := gobject.GBool(visible)
	defer b.Free()
	C.gtk_widget_set_visible(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Widget) GetHalign() int {
	return int(C.gtk_widget_get_halign(self.object))
}

func (self *Widget) SetHalign(gtk_Align int) {
	C.gtk_widget_set_halign(self.object, C.GtkAlign(gtk_Align))
}

func (self *Widget) GetValign() int {
	return int(C.gtk_widget_get_valign(self.object))
}

func (self *Widget) SetValign(gtk_align int) {
	C.gtk_widget_set_valign(self.object, C.GtkAlign(gtk_align))
}

func (self *Widget) GetCanDefault() bool {
	b := C.gtk_widget_get_can_default(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) SetCanDefault(canDefault bool) {
	b := gobject.GBool(canDefault)
	defer b.Free()
	C.gtk_widget_set_can_default(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Widget) GetPreferredSize() (minimumSize, naturalSize Requisition) {
	var min, nat C.GtkRequisition
	C.gtk_widget_get_preferred_size(self.object, &min, &nat)

	minimumSize = Requisition{int(min.width), int(min.height)}
	naturalSize = Requisition{int(nat.width), int(nat.height)}
	return
}

//////////////////////////////
// END GtkWidget
////////////////////////////// }}}

// GtkContainer {{{
//////////////////////////////

// Container Type
type Container struct {
	object *C.GtkContainer
	*Widget
}

// New Container from widget
func NewContainer(o unsafe.Pointer) *Container {
	con := &Container{}
	con.object = C.to_GtkContainer(o)

	if gobject.IsObjectFloating(con) {
		gobject.RefSink(con)
	} else {
		gobject.Ref(con)
	}
	con.Widget = NewWidget(o)
	containerFinalizer(con)

	return con
}

// Finalizer for Container struct
func containerFinalizer(c *Container) {
	runtime.SetFinalizer(c, func(c *Container) { gobject.Unref(c) })
}

// To be Object-like
func (self Container) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Container) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Container) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Container) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Widget-like
func (self Container) W() *Widget {
	return self.Widget
}

// interfaces
func (self *Container) Add(w WidgetLike) {
	C.gtk_container_add(self.object, w.W().object)
}

func (self *Container) Remove(w WidgetLike) {
	C.gtk_container_remove(self.object, w.W().object)
}

func (self *Container) GetResizeMode() int {
	return int(C.gtk_container_get_resize_mode(self.object))
}

func (self *Container) SetResizeMode(gtk_ResizeMode int) {
	C.gtk_container_set_resize_mode(self.object, C.GtkResizeMode(gtk_ResizeMode))
}

func (self *Container) CheckResize() {
	C.gtk_container_check_resize(self.object)
}

func (self *Container) ForEach(f interface{}, data ...interface{}) {
	forClosure, id := gobject.CreateCustomClosure(f, data...)

	if forClosure == nil {
		return
	}

	addToClosures(id, forClosure)

	var cId C.gint64 = C.gint64(id)
	C._gtk_container_foreach(self.object, &cId)

	removeFromClosures(id)
}

//TODO: gtk_container_get_children
//TODO: gtk_container_get_path_for_child

func (self *Container) SetReallocateRedraws(needRedraws bool) {
	b := gobject.GBool(needRedraws)
	defer b.Free()
	C.gtk_container_set_reallocate_redraws(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Container) GetFocusChild() WidgetLike {
	w := C.gtk_container_get_focus_child(self.object)
	i, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return i.(WidgetLike)
	}
	return nil
}

func (self *Container) GetFocusVadjustment() *Adjustment {
	a := C.gtk_container_get_focus_vadjustment(self.object)
	if a == nil {
		return nil
	}
	adj, err := gobject.ConvertToGo(unsafe.Pointer(a))
	if err == nil {
		return adj.(*Adjustment)
	}
	return nil
}

func (self *Container) SetFocusVadjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_vadjustment(self.object, adjustment.object)
}

func (self *Container) GetFocusHadjustment() *Adjustment {
	a := C.gtk_container_get_focus_hadjustment(self.object)
	if a == nil {
		return nil
	}
	adj, err := gobject.ConvertToGo(unsafe.Pointer(a))
	if err == nil {
		return adj.(*Adjustment)
	}
	return nil
}

func (self *Container) SetFocusHadjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_hadjustment(self.object, adjustment.object)
}

func (self *Container) GetBorderWidth() uint {
	return uint(C.gtk_container_get_border_width(self.object))
}

func (self *Container) SetBorderWidth(width uint) {
	C.gtk_container_set_border_width(self.object, C.guint(width))
}

//////////////////////////////
// END GtkContainer
////////////////////////////// }}}

// GtkBin {{{
//////////////////////////////

// GtkBin type
type Bin struct {
	object *C.GtkBin
	*Container
}

func NewBin(obj unsafe.Pointer) *Bin {
	b := &Bin{}
	b.object = C.to_GtkBin(obj)

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	} else {
		gobject.Ref(b)
	}
	b.Container = NewContainer(obj)
	binFinalizer(b)

	return b
}

// Clear Bin struct when it goes out of reach
func binFinalizer(b *Bin) {
	runtime.SetFinalizer(b, func(b *Bin) { gobject.Unref(b) })
}

// Conversion funcs
func newBinFromNative(obj unsafe.Pointer) interface{} {
	return NewBin(obj)
}

func nativeFromBin(b interface{}) *gobject.GValue {
	bin, ok := b.(*Bin)
	if ok {
		gv := gobject.CreateCGValue(GtkType.BIN, bin.ToNative())
		return gv
	}
	return nil
}

// To be Object-like
func (self Bin) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Bin) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Bin) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Bin) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-lie
func (self Bin) C() *Container {
	return self.Container
}

// Bin interface

func (self *Bin) GetChild() WidgetLike {
	ch := C.gtk_bin_get_child(self.object)
	if ch == nil {
		return nil
	}
	w, err := gobject.ConvertToGo(unsafe.Pointer(ch))
	if err == nil {
		return w.(WidgetLike)
	}
	return nil
}

//////////////////////////////
// End GtkBin
////////////////////////////// }}}

// GtkRange {{{
//////////////////////////////

// GtkRange type
type Range struct {
	object *C.GtkRange
	*Widget
}

// Clear Range struct when it goes out of reach
func rangeFinalizer(r *Range) {
	runtime.SetFinalizer(r, func(r *Range) { gobject.Unref(r) })
}

// Conversion functions for gobject registration map
func newRangeFromNative(obj unsafe.Pointer) interface{} {
	r := &Range{}
	r.object = C.to_GtkRange(obj)

	if gobject.IsObjectFloating(r) {
		gobject.RefSink(r)
	} else {
		gobject.Ref(r)
	}
	r.Widget = NewWidget(obj)
	rangeFinalizer(r)

	return r
}

func nativeFromRange(r interface{}) *gobject.GValue {
	ran, ok := r.(*Range)
	if ok {
		gv := gobject.CreateCGValue(GtkType.RANGE, ran.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Range) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Range) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Range) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Range) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Range) W() *Widget {
	return self.Widget
}

// Range interface
func (self *Range) GetFillLevel() float64 {
	return float64(C.gtk_range_get_fill_level(self.object))
}

func (self *Range) GetRestrictToFillLevel() bool {
	b := C.gtk_range_get_restrict_to_fill_level(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Range) GetShowFillLevel() bool {
	b := C.gtk_range_get_show_fill_level(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Range) SetFillLevel(fillLevel float64) {
	C.gtk_range_set_fill_level(self.object, C.gdouble(fillLevel))
}

func (self *Range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	b := gobject.GBool(restrictToFillLevel)
	defer b.Free()
	C.gtk_range_set_restrict_to_fill_level(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Range) SetShowFillLevel(showFillLevel bool) {
	b := gobject.GBool(showFillLevel)
	defer b.Free()
	C.gtk_range_set_show_fill_level(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Range) GetAdjustment() *Adjustment {
	a := C.gtk_range_get_adjustment(self.object)
	adj, err := gobject.ConvertToGo(unsafe.Pointer(&a))
	if err == nil {
		return adj.(*Adjustment)
	}
	return nil
}

func (self *Range) SetAdjustment(adjustment *Adjustment) {
	C.gtk_range_set_adjustment(self.object, adjustment.object)
}

func (self *Range) GetInverted() bool {
	b := C.gtk_range_get_inverted(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Range) SetInverted(inverted bool) {
	b := gobject.GBool(inverted)
	C.gtk_range_set_inverted(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Range) GetValue() float64 {
	return float64(C.gtk_range_get_value(self.object))
}

func (self *Range) SetValue(value float64) {
	C.gtk_range_set_value(self.object, C.gdouble(value))
}

func (self *Range) SetIncrements(step, page float64) {
	C.gtk_range_set_increments(self.object, C.gdouble(step), C.gdouble(page))
}

func (self *Range) SetRange(min, max float64) {
	C.gtk_range_set_range(self.object, C.gdouble(min), C.gdouble(max))
}

func (self *Range) GetRoundDigits() int {
	return int(C.gtk_range_get_round_digits(self.object))
}

func (self *Range) SetRoundDigits(roundDigits int) {
	C.gtk_range_set_round_digits(self.object, C.gint(roundDigits))
}

func (self *Range) SetLowerStepperSensitivity(gtk_sensitivityType int) {
	C.gtk_range_set_lower_stepper_sensitivity(self.object, C.GtkSensitivityType(gtk_sensitivityType))
}

func (self *Range) GetLowerStepperSensitivity() int {
	return int(C.gtk_range_get_lower_stepper_sensitivity(self.object))
}

func (self *Range) SetUpperStepperSensitivity(gtk_sensitivityType int) {
	C.gtk_range_set_upper_stepper_sensitivity(self.object, C.GtkSensitivityType(gtk_sensitivityType))
}

func (self *Range) GetUpperStepperSensitivity() int {
	return int(C.gtk_range_get_upper_stepper_sensitivity(self.object))
}

func (self *Range) GetFlippable() bool {
	b := C.gtk_range_get_flippable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Range) SetFlippable(flippable bool) {
	b := gobject.GBool(flippable)
	defer b.Free()
	C.gtk_range_set_flippable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Range) GetMinSliderSize() int {
	return int(C.gtk_range_get_min_slider_size(self.object))
}

//TODO: gtk_range_get_range_rect

func (self *Range) GetSliderRange() (start, end int) {
	var s C.gint
	var e C.gint
	C.gtk_range_get_slider_range(self.object, &s, &e)
	start = int(s)
	end = int(e)
	return
}

func (self *Range) GetSliderSizeFixed() bool {
	b := C.gtk_range_get_slider_size_fixed(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Range) SetMinSliderSize(minSize int) {
	C.gtk_range_set_min_slider_size(self.object, C.gint(minSize))
}

func (self *Range) SetSliderSizeFixed(sizeFixed bool) {
	b := gobject.GBool(sizeFixed)
	defer b.Free()
	C.gtk_range_set_slider_size_fixed(self.object, *((*C.gboolean)(b.GetPtr())))
}

//////////////////////////////
// END GtkRange
////////////////////////////// }}}

// End Base Structs }}}

// Windows {{{

// GtkWindow {{{
//////////////////////////////

// Window type
type Window struct {
	object *C.GtkWindow
	*Bin
}

// Create new window
func NewWindow(windowType int) (w *Window) {
	w = &Window{}
	o := C.gtk_window_new(C.GtkWindowType(windowType))
	w.object = C.to_GtkWindow(unsafe.Pointer(o))
	gobject.RefSink(w)

	w.Bin = NewBin(unsafe.Pointer(o))
	windowFinalizer(w)

	return w
}

// Clear Window struct when it goes out of reach
func windowFinalizer(w *Window) {
	runtime.SetFinalizer(w, func(w *Window) { gobject.Unref(w) })
}

// Conversion function for gobject registration map
func newWindowFromNative(obj unsafe.Pointer) interface{} {
	w := &Window{}
	w.object = C.to_GtkWindow(obj)

	if gobject.IsObjectFloating(w) {
		gobject.RefSink(w)
	} else {
		gobject.Ref(w)
	}
	w.Bin = NewBin(obj)
	windowFinalizer(w)

	return w
}

func nativeFromWindow(w interface{}) *gobject.GValue {
	win, ok := w.(*Window)
	if ok {
		gv := gobject.CreateCGValue(GtkType.WINDOW, win.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Window) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Window) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Window) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Window) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be bin-like
func (self Window) CBin() *Bin {
	return self.Bin
}

// Window interface

func (self *Window) AddAccelGroup(accelGroup *AccelGroup) {
	C.gtk_window_add_accel_group(self.object, accelGroup.object)
}

func (self *Window) RemoveAccelGroup(accelGroup *AccelGroup) {
	C.gtk_window_remove_accel_group(self.object, accelGroup.object)
}

func (self *Window) SetTitle(title string) {
	gobject.SetProperty(self, "title", title)
}

func (self *Window) GetTitle() string {
	t := gobject.GetProperty(self, "title")
	defer t.Free()
	return gobject.GoString(t.GetPtr())
}

func (self *Window) SetWMClass(name, class string) {
	n := gobject.GString(name)
	c := gobject.GString(class)
	defer n.Free()
	defer c.Free()
	C.gtk_window_set_wmclass(self.object, (*C.gchar)(n.GetPtr()),
		(*C.gchar)(c.GetPtr()))
}

func (self *Window) SetStartupId(id string) {
	gobject.SetProperty(self, "startup-id", id)
}

func (self *Window) ActivateFocus() bool {
	b := C.gtk_window_activate_focus(self.object)
	gb, _ := gobject.ConvertToGo(unsafe.Pointer(&b), gobject.G_TYPE_BOOLEAN)
	return gb.(bool)
}

func (self *Window) SetFocus(w WidgetLike) {
	C.gtk_window_set_focus(self.object, w.W().object)
}

func (self *Window) SetDefault(w WidgetLike) {
	C.gtk_window_set_default(self.object, w.W().object)
}

func (self *Window) ActivateDefault() bool {
	b := C.gtk_window_activate_default(self.object)
	gb, _ := gobject.ConvertToGo(unsafe.Pointer(&b), gobject.G_TYPE_BOOLEAN)
	return gb.(bool)
}

func (self *Window) SetTransientFor(parent *Window) {
	C.gtk_window_set_transient_for(self.object, parent.object)
}

func (self *Window) GetTransientFor() *Window {
	w := unsafe.Pointer(C.gtk_window_get_transient_for(self.object))
	if w == nil {
		return nil
	}
	w1 := newWindowFromNative(w).(Window)
	return &w1
}

func (self *Window) SetOpacity(opacity float64) {
	gobject.SetProperty(self, "opacity", opacity)
}

func (self *Window) GetOpacity() float64 {
	return float64(C.gtk_window_get_opacity(self.object))
}

func (self *Window) SetSkipTaskbarHint(setting bool) {
	gobject.SetProperty(self, "skip-taskbar-hint", setting)
}

func (self *Window) GetSkipTaskbarHint() bool {
	b := gobject.GetProperty(self, "skip-taskbar-hint")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) GetFocus() WidgetLike {
	widget := C.gtk_window_get_focus(self.object)
	val, err := gobject.ConvertToGo(unsafe.Pointer(widget))

	if err != nil {
		return nil
	}
	return val.(WidgetLike)
}

func (self *Window) GetDefaultWidget() WidgetLike {
	widget := C.gtk_window_get_default_widget(self.object)
	val, err := gobject.ConvertToGo(unsafe.Pointer(widget))

	if err != nil {
		return nil
	}

	return val.(WidgetLike)
}

func (self *Window) SetSkipPagerHint(setting bool) {
	gobject.SetProperty(self, "skip-pager-hint", setting)
}

func (self *Window) GetSkipPagerHint() bool {
	b := gobject.GetProperty(self, "skip-pager-hint")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetUrgencyHint(setting bool) {
	gobject.SetProperty(self, "urgency-hint", setting)
}

func (self *Window) GetUrgencyHint() bool {
	b := gobject.GetProperty(self, "urgency-hint")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetAcceptFocus(setting bool) {
	gobject.SetProperty(self, "accept-focus", setting)
}

func (self *Window) GetAcceptFocus() bool {
	b := gobject.GetProperty(self, "accept-focus")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetFocusOnMap(setting bool) {
	gobject.SetProperty(self, "focus-on-map", setting)
}

func (self *Window) GetFocusOnMap() bool {
	b := gobject.GetProperty(self, "focus-on-map")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetDestroyWithParent(setting bool) {
	gobject.SetProperty(self, "destroy-with-parent", setting)
}

func (self *Window) GetDestroyWithParent() bool {
	b := gobject.GetProperty(self, "destroy-with-parent")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetMnemonicsVisible(setting bool) {
	gobject.SetProperty(self, "mnemonics-visible", setting)
}

func (self *Window) GetMnemonicsVisible() bool {
	b := gobject.GetProperty(self, "mnemonics-visible")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetResizable(setting bool) {
	gobject.SetProperty(self, "resizable", setting)
}

func (self *Window) GetResizable() bool {
	b := gobject.GetProperty(self, "resizable")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) IsActive() bool {
	b := C.gtk_window_is_active(self.object)
	cb, _ := gobject.ConvertToGo(unsafe.Pointer(&b), gobject.G_TYPE_BOOLEAN)
	return cb.(bool)
}

func (self *Window) SetScreen(screen *gdk3.Screen) {
	if screen != nil {
		C.gtk_window_set_screen(self.object, (*C.GdkScreen)(screen.ToNative()))
	}
}

func (self *Window) HasTopLevelFocus() bool {
	b := C.gtk_window_has_toplevel_focus(self.object)
	cb, _ := gobject.ConvertToGo(unsafe.Pointer(&b), gobject.G_TYPE_BOOLEAN)
	return cb.(bool)
}

func (self *Window) SetDecorated(setting bool) {
	gobject.SetProperty(self, "decorated", setting)
}

func (self *Window) GetDecorated() bool {
	b := gobject.GetProperty(self, "decorated")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetDeletable(setting bool) {
	gobject.SetProperty(self, "deletable", setting)
}

func (self *Window) GetDeletable() bool {
	b := gobject.GetProperty(self, "deletable")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetModal(setting bool) {
	gobject.SetProperty(self, "modal", setting)
}

func (self *Window) GetModal() bool {
	b := gobject.GetProperty(self, "modal")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetDefaultSize(width, height int) {
	C.gtk_window_set_default_size(self.object, C.gint(width), C.gint(height))
}

func (self *Window) WindowPresent() {
	C.gtk_window_present(self.object)
}

func (self *Window) Iconify() {
	C.gtk_window_iconify(self.object)
}

func (self *Window) Deiconify() {
	C.gtk_window_deiconify(self.object)
}

func (self *Window) Stick() {
	C.gtk_window_stick(self.object)
}

func (self *Window) UnStick() {
	C.gtk_window_unstick(self.object)
}

func (self *Window) Maximize() {
	C.gtk_window_maximize(self.object)
}

func (self *Window) UnMaximize() {
	C.gtk_window_unmaximize(self.object)
}

func (self *Window) FullScreen() {
	C.gtk_window_fullscreen(self.object)
}

func (self *Window) UnFullScreen() {
	C.gtk_window_unfullscreen(self.object)
}

func (self *Window) SetKeepAbove(setting bool) {
	b := gobject.ConvertToC(setting)
	defer b.Free()
	C.gtk_window_set_keep_above(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Window) SetKeepBelow(setting bool) {
	b := gobject.ConvertToC(setting)
	defer b.Free()
	C.gtk_window_set_keep_below(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Window) SetApplication(app *Application) {
	gobject.SetProperty(self, "application", app)
}

func (self *Window) GetApplication() *Application {
	a := gobject.GetProperty(self, "application")
	defer a.Free()
	app, _ := gobject.ConvertToGo(a.GetPtr())
	if app != nil {
		return app.(*Application)
	}

	return nil
}

//////////////////////////////
// END GtkWindow
////////////////////////////// }}}

// GtkDialog {{{
//////////////////////////////

// GtkDialog type
type Dialog struct {
	object *C.GtkDialog
	*Window
}

func NewDialog() *Dialog {
	d := &Dialog{}
	o := C.gtk_dialog_new()
	d.object = C.to_GtkDialog(unsafe.Pointer(o))

	if gobject.IsObjectFloating(d) {
		gobject.RefSink(d)
	}
	d.Window = newWindowFromNative(unsafe.Pointer(o)).(*Window)
	dialogFinalizer(d)

	return d
}

func NewDialogWithButtons(title string, parent *Window, flags int, butAndID B) *Dialog {
	// Must have at least one button
	if len(butAndID) == 0 {
		return nil
	}

	t := gobject.GString(title)
	defer t.Free()

	firstButton := butAndID[0].Text
	fb := gobject.GString(firstButton)
	defer fb.Free()

	firstId := butAndID[0].Response

	var wparent *C.GtkWindow = nil
	if parent != nil {
		wparent = parent.object
	}

	d := &Dialog{}
	o := C._dialog_new_with_buttons((*C.gchar)(t.GetPtr()), wparent, C.GtkDialogFlags(flags),
		(*C.gchar)(fb.GetPtr()), C.gint(firstId))
	d.object = C.to_GtkDialog(unsafe.Pointer(o))

	if gobject.IsObjectFloating(d) {
		gobject.RefSink(d)
	}
	d.Window = newWindowFromNative(unsafe.Pointer(o)).(*Window)
	dialogFinalizer(d)
	d.AddButtons(butAndID[1:])

	return d
}

// Clear Dialog struct when it goes out of reach
func dialogFinalizer(d *Dialog) {
	runtime.SetFinalizer(d, func(d *Dialog) { gobject.Unref(d) })
}

// Conversion function for gobject registration map
func newDialogFromNative(obj unsafe.Pointer) interface{} {
	d := &Dialog{}
	d.object = C.to_GtkDialog(obj)

	if gobject.IsObjectFloating(d) {
		gobject.RefSink(d)
	} else {
		gobject.Ref(d)
	}
	d.Window = newWindowFromNative(obj).(*Window)
	dialogFinalizer(d)

	return d
}

func nativeFromDialog(d interface{}) *gobject.GValue {
	dialog, ok := d.(*Dialog)
	if ok {
		gv := gobject.CreateCGValue(GtkType.DIALOG, dialog.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self Dialog) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Dialog) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Dialog) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Dialog) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be window-like
func (self Dialog) Wnd() *Window {
	return self.Window
}

// Dialog interface
func (self *Dialog) Run() int {
	response := C.gtk_dialog_run(self.object)
	return int(response)
}

func (self *Dialog) Response(responseId int) {
	C.gtk_dialog_response(self.object, C.gint(responseId))
}

func (self *Dialog) AddButton(buttonText string, responseId int) *Button {
	s := gobject.GString(buttonText)
	defer s.Free()
	b := C.gtk_dialog_add_button(self.object, (*C.gchar)(s.GetPtr()), C.gint(responseId))
	btn, err := gobject.ConvertToGo(unsafe.Pointer(&b))
	if err == nil {
		return btn.(*Button)
	}
	return nil
}

func (self *Dialog) AddButtons(buttonAndId B) {
	for _, btn := range buttonAndId {
		self.AddButton(btn.Text, btn.Response)
	}
}

func (self *Dialog) AddActionWidget(w WidgetLike, responseId int) {
	C.gtk_dialog_add_action_widget(self.object, w.W().object, C.gint(responseId))
}

func (self *Dialog) SetDefaultResponse(responseId int) {
	C.gtk_dialog_set_default_response(self.object, C.gint(responseId))
}

func (self *Dialog) SetResponseSensitive(responseId int, setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_dialog_set_response_sensitive(self.object, C.gint(responseId), *((*C.gboolean)(b.GetPtr())))
}

func (self *Dialog) GetResponseForWidget(w WidgetLike) int {
	i := C.gtk_dialog_get_response_for_widget(self.object, w.W().object)
	return int(i)
}

func (self *Dialog) GetWidgetForResponse(responseId int) WidgetLike {
	cw := C.gtk_dialog_get_widget_for_response(self.object, C.gint(responseId))
	if cw != nil {
		w, err := gobject.ConvertToGo(unsafe.Pointer(cw))
		if err == nil {
			return w.(WidgetLike)
		}
	}
	return nil
}

func (self *Dialog) GetActionArea() WidgetLike {
	cw := C.gtk_dialog_get_action_area(self.object)
	if cw != nil {
		w, err := gobject.ConvertToGo(unsafe.Pointer(cw))
		if err == nil {
			return w.(WidgetLike)
		}
	}
	return nil
}

func (self *Dialog) GetContentArea() WidgetLike {
	cw := C.gtk_dialog_get_content_area(self.object)
	if cw != nil {
		w, err := gobject.ConvertToGo(unsafe.Pointer(cw))
		if err == nil {
			return w.(WidgetLike)
		}
	}
	return nil
}

//////////////////////////////
// END GtkDialog
////////////////////////////// }}}

// GtkMessageDialog {{{
//////////////////////////////

// GtkMessageDialog type
type MessageDialog struct {
	object *C.GtkMessageDialog
	*Dialog
}

func NewMessageDialog(parent *Window, flags int, mtype int, buttons int, messageFormat string, argsToMessage ...interface{}) *MessageDialog {
	mdiag := &MessageDialog{}
	var message string = messageFormat

	// If additional args, then call Sprintf with those args
	// and save formated message
	if len(argsToMessage) > 0 {
		message = fmt.Sprintf(messageFormat, argsToMessage...)
	}

	cmessage := gobject.GString(message)
	defer cmessage.Free()

	// If parent is not nil we will feed the parent object
	var pwin *C.GtkWindow = nil
	if parent != nil {
		pwin = parent.object
	}

	o := C._new_message_dialog(pwin, C.GtkDialogFlags(flags), C.GtkMessageType(mtype), C.GtkButtonsType(buttons),
		(*C.gchar)(cmessage.GetPtr()))
	mdiag.object = C.to_GtkMessageDialog(unsafe.Pointer(o))
	if gobject.IsObjectFloating(mdiag) {
		gobject.RefSink(mdiag)
	}
	mdiag.Dialog = newDialogFromNative(unsafe.Pointer(o)).(*Dialog)
	messageDialogFinalizer(mdiag)

	return mdiag
}

// Clear MessageDialog struct when it goes out of reach
func messageDialogFinalizer(d *MessageDialog) {
	runtime.SetFinalizer(d, func(d *MessageDialog) { gobject.Unref(d) })
}

// Conversion function for gobject registration map
func newMessageDialogFromNative(obj unsafe.Pointer) interface{} {
	d := &MessageDialog{}
	d.object = C.to_GtkMessageDialog(obj)

	if gobject.IsObjectFloating(d) {
		gobject.RefSink(d)
	} else {
		gobject.Ref(d)
	}
	d.Dialog = newDialogFromNative(obj).(*Dialog)
	messageDialogFinalizer(d)
	return d
}

func nativeFromMessageDialog(d interface{}) *gobject.GValue {
	dialog, ok := d.(*MessageDialog)
	if ok {
		gv := gobject.CreateCGValue(GtkType.MESSAGE_DIALOG, dialog.ToNative())
		return gv
	}

	return nil
}

// To be object like
func (self MessageDialog) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self MessageDialog) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self MessageDialog) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self MessageDialog) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be dialog-like
func (self MessageDialog) D() *Dialog {
	return self.Dialog
}

// Message dialog interface
func (self *MessageDialog) SetMarkup(markup string) {
	m := gobject.GString(markup)
	defer m.Free()
	C.gtk_message_dialog_set_markup(self.object, (*C.gchar)(m.GetPtr()))
}

func (self *MessageDialog) SetImage(w WidgetLike) {
	C.gtk_message_dialog_set_image(self.object, w.W().object)
}

func (self *MessageDialog) GetImage() WidgetLike {
	im := C.gtk_message_dialog_get_image(self.object)
	if im != nil {
		i, err := gobject.ConvertToGo(unsafe.Pointer(im))
		if err == nil {
			return i.(WidgetLike)
		}
	}
	return nil
}

func (self *MessageDialog) GetMessageArea() WidgetLike {
	ma := C.gtk_message_dialog_get_message_area(self.object)
	if ma != nil {
		m, err := gobject.ConvertToGo(unsafe.Pointer(ma))
		if err == nil {
			return m.(WidgetLike)
		}
	}
	return nil
}

func (self *MessageDialog) FormatSecondaryText(messageFormat string, args ...interface{}) {
	var message string = messageFormat
	if len(args) > 0 {
		message = fmt.Sprintf(messageFormat, args...)
	}

	cmessage := gobject.GString(message)
	defer cmessage.Free()

	C._gtk_message_dialog_format_secondary_text(self.object, (*C.gchar)(cmessage.GetPtr()))
}

func (self *MessageDialog) FormatSecondaryMarkup(format string, args ...interface{}) {
	var message string = format
	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	}
	cmessage := gobject.GString(message)
	defer cmessage.Free()

	C._gtk_message_dialog_format_secondary_markup(self.object, (*C.gchar)(cmessage.GetPtr()))
}

//////////////////////////////
// END GtkMessageDialog
////////////////////////////// }}}

// GtkInvisible {{{
//////////////////////////////

// GtkInvisible type
type Invisible struct {
	object *C.GtkInvisible
	*Widget
}

func NewInvisible() *Invisible {
	i := &Invisible{}
	o := C.gtk_invisible_new()
	i.object = C.to_GtkInvisible(unsafe.Pointer(o))

	if gobject.IsObjectFloating(i) {
		gobject.RefSink(i)
	}
	i.Widget = NewWidget(unsafe.Pointer(o))
	invisibleFinalizer(i)

	return i
}

func NewInvisibleForScreen(screen *gdk3.Screen) *Invisible {
	i := &Invisible{}
	o := C.gtk_invisible_new_for_screen((*C.GdkScreen)(screen.ToNative()))
	i.object = C.to_GtkInvisible(unsafe.Pointer(o))

	if gobject.IsObjectFloating(i) {
		gobject.RefSink(i)
	}
	i.Widget = NewWidget(unsafe.Pointer(o))
	invisibleFinalizer(i)

	return i
}

// Clear Invisible struct when it goes out of reach
func invisibleFinalizer(i *Invisible) {
	runtime.SetFinalizer(i, func(i *Invisible) { gobject.Unref(i) })
}

// Conversion funcs
func newInvisibleFromNative(obj unsafe.Pointer) interface{} {
	i := &Invisible{}
	i.object = C.to_GtkInvisible(obj)

	if gobject.IsObjectFloating(i) {
		gobject.RefSink(i)
	} else {
		gobject.Ref(i)
	}
	i.Widget = NewWidget(obj)
	invisibleFinalizer(i)

	return i
}

func nativeFromInvisible(i interface{}) *gobject.GValue {
	inv, ok := i.(*Invisible)
	if ok {
		gv := gobject.CreateCGValue(GtkType.INVISIBLE, inv.ToNative())
		return gv
	}
	return nil
}

// To be object like
func (self Invisible) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Invisible) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Invisible) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Invisible) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Invisible) W() *Widget {
	return self.Widget
}

// Invisible interface

func (self *Invisible) SetScreen(screen *gdk3.Screen) {
	C.gtk_invisible_set_screen(self.object, (*C.GdkScreen)(screen.ToNative()))
}

func (self *Invisible) GetScreen() *gdk3.Screen {
	s := C.gtk_invisible_get_screen(self.object)

	if scr, err := gobject.ConvertToGo(unsafe.Pointer(s)); err == nil {
		return scr.(*gdk3.Screen)
	}
	return nil
}

//////////////////////////////
// END GtkInvisible
////////////////////////////// }}}

// GtkAssistant {{{
//////////////////////////////

// GtkAssistant type
type Assistant struct {
	object *C.GtkAssistant
	*Window
}

// Clear Assistant struct when it goes out of reach
func assistantFinalizer(as *Assistant) {
	runtime.SetFinalizer(as, func(as *Assistant) { gobject.Unref(as) })
}

func NewAssistant() *Assistant {
	as := &Assistant{}
	o := C.gtk_assistant_new()
	as.object = C.to_GtkAssistant(unsafe.Pointer(o))

	if gobject.IsObjectFloating(as) {
		gobject.RefSink(as)
	}
	as.Window = newWindowFromNative(unsafe.Pointer(o)).(*Window)
	assistantFinalizer(as)

	return as
}

// Conversion funcs
func newAssistantFromNative(obj unsafe.Pointer) interface{} {
	as := &Assistant{}
	as.object = C.to_GtkAssistant(obj)

	if gobject.IsObjectFloating(as) {
		gobject.RefSink(as)
	} else {
		gobject.Ref(as)
	}
	as.Window = newWindowFromNative(obj).(*Window)
	assistantFinalizer(as)

	return as
}

func nativeFromAssistant(as interface{}) *gobject.GValue {
	assistant, ok := as.(*Assistant)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ASSISTANT, assistant.ToNative())
		return gv
	}
	return nil
}

// To be object like
func (self Assistant) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Assistant) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Assistant) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Assistant) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Assistant interface

func (self *Assistant) GetCurrentPage() int {
	return int(C.gtk_assistant_get_current_page(self.object))
}

func (self *Assistant) SetCurrentPage(pageNum int) {
	C.gtk_assistant_set_current_page(self.object, C.gint(pageNum))
}

func (self *Assistant) GetNPages() int {
	return int(C.gtk_assistant_get_n_pages(self.object))
}

func (self *Assistant) GetNthPage(pageNum int) WidgetLike {
	w := C.gtk_assistant_get_nth_page(self.object, C.gint(pageNum))

	if w == nil {
		return nil
	}

	wid, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

func (self *Assistant) PrependPage(page WidgetLike) int {
	return int(C.gtk_assistant_prepend_page(self.object, page.W().object))
}

func (self *Assistant) AppendPage(page WidgetLike) int {
	return int(C.gtk_assistant_append_page(self.object, page.W().object))
}

func (self *Assistant) InsertPage(page WidgetLike, position int) int {
	return int(C.gtk_assistant_insert_page(self.object, page.W().object, C.gint(position)))
}

func (self *Assistant) RemovePage(pageNum int) {
	C.gtk_assistant_remove_page(self.object, C.gint(pageNum))
}

//TODO: gtk_assistant_set_forward_page_func

func (self *Assistant) SetPageType(page WidgetLike, gtk_AssistantPage int) {
	C.gtk_assistant_set_page_type(self.object, page.W().object, C.GtkAssistantPageType(gtk_AssistantPage))
}

func (self *Assistant) GetPageType(page WidgetLike) int {
	return int(C.gtk_assistant_get_page_type(self.object, page.W().object))
}

func (self *Assistant) SetPageTitle(page WidgetLike, title string) {
	s := gobject.GString(title)
	defer s.Free()

	C.gtk_assistant_set_page_title(self.object, page.W().object, (*C.gchar)(s.GetPtr()))
}

func (self *Assistant) GetPageTitle(page WidgetLike) string {
	s := C.gtk_assistant_get_page_title(self.object, page.W().object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Assistant) SetPageComplete(page WidgetLike, complete bool) {
	b := gobject.GBool(complete)
	defer b.Free()
	C.gtk_assistant_set_page_complete(self.object, page.W().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Assistant) GetPageComplete(page WidgetLike) bool {
	b := C.gtk_assistant_get_page_complete(self.object, page.W().object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Assistant) AddActionWidget(child WidgetLike) {
	C.gtk_assistant_add_action_widget(self.object, child.W().object)
}

func (self *Assistant) RemoveActionWidget(child WidgetLike) {
	C.gtk_assistant_remove_action_widget(self.object, child.W().object)
}

func (self *Assistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(self.object)
}

func (self *Assistant) Commit() {
	C.gtk_assistant_commit(self.object)
}

func (self *Assistant) NextPage() {
	C.gtk_assistant_next_page(self.object)
}

func (self *Assistant) PreviousPage() {
	C.gtk_assistant_previous_page(self.object)
}

//////////////////////////////
// END GtkAssistant
////////////////////////////// }}}

// End Windows }}}

// Display Widgets {{{

// GtkLabel {{{
//////////////////////////////

// GtkLabel type
type Label struct {
	object *C.GtkLabel
	*Widget
}

func NewLabel(str string) *Label {
	l := &Label{}

	s := gobject.GString(str)
	defer s.Free()

	o := C.gtk_label_new((*C.gchar)(s.GetPtr()))
	l.object = C.to_GtkLabel(unsafe.Pointer(o))

	if gobject.IsObjectFloating(l) {
		gobject.RefSink(l)
	}
	l.Widget = NewWidget(unsafe.Pointer(o))
	labelFinalizer(l)

	return l
}

func NewLabelWithMnemonic(str string) *Label {
	l := &Label{}

	s := gobject.GString(str)
	defer s.Free()

	o := C.gtk_label_new_with_mnemonic((*C.gchar)(s.GetPtr()))
	l.object = C.to_GtkLabel(unsafe.Pointer(o))

	if gobject.IsObjectFloating(l) {
		gobject.RefSink(l)
	}
	l.Widget = NewWidget(unsafe.Pointer(o))
	labelFinalizer(l)

	return l
}

// Clear label struct when it goes out of reach
func labelFinalizer(l *Label) {
	runtime.SetFinalizer(l, func(l *Label) { gobject.Unref(l) })
}

// Conversion function for gobject registration map
func newLabelFromNative(obj unsafe.Pointer) interface{} {
	l := &Label{}
	l.object = C.to_GtkLabel(obj)

	if gobject.IsObjectFloating(l) {
		gobject.RefSink(l)
	} else {
		gobject.Ref(l)
	}
	l.Widget = NewWidget(unsafe.Pointer(l.object))
	labelFinalizer(l)

	return l
}

func nativeFromLabel(label interface{}) *gobject.GValue {
	l, ok := label.(*Label)
	if ok {
		gv := gobject.CreateCGValue(GtkType.LABEL, l.ToNative())
		return gv
	}

	return nil
}

// To be object like
func (self Label) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Label) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Label) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self Label) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Label) W() *Widget {
	return self.Widget
}

// Label interface

func (self *Label) SetText(str string) {
	s := gobject.GString(str)
	defer s.Free()
	C.gtk_label_set_text(self.object, (*C.gchar)(s.GetPtr()))
}

//TODO: gtk_label_set_attributes

func (self *Label) SetMarkup(str string) {
	s := gobject.GString(str)
	defer s.Free()
	C.gtk_label_set_markup(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Label) SetMarkupWithMnemonic(str string) {
	s := gobject.GString(str)
	defer s.Free()
	C.gtk_label_set_markup_with_mnemonic(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Label) SetPattern(str string) {
	s := gobject.GString(str)
	defer s.Free()
	C.gtk_label_set_pattern(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Label) SetJustify(jtype int) {
	C.gtk_label_set_justify(self.object, C.GtkJustification(jtype))
}

//TODO: gtk_label_set_ellipsize

func (self *Label) SetWidthChars(n_chars int) {
	C.gtk_label_set_width_chars(self.object, C.gint(n_chars))
}

func (self *Label) SetMaxWidthChars(n_chars int) {
	C.gtk_label_set_max_width_chars(self.object, C.gint(n_chars))
}

func (self *Label) SetLineWrap(wrap bool) {
	b := gobject.GBool(wrap)
	defer b.Free()
	C.gtk_label_set_line_wrap(self.object, *((*C.gboolean)(b.GetPtr())))
}

//TODO: gtk_label_set_line_wrap_mode

func (self *Label) GetLayoutOffsets() (x, y int) {
	var cx C.gint
	var cy C.gint
	C.gtk_label_get_layout_offsets(self.object, &cx, &cy)
	return int(cx), int(cy)
}

func (self *Label) GetMnemonicKeyVal() uint {
	ckv := C.gtk_label_get_mnemonic_keyval(self.object)
	return uint(ckv)
}

func (self *Label) GetSelectable() bool {
	sel := C.gtk_label_get_selectable(self.object)
	return gobject.GoBool(unsafe.Pointer(&sel))
}

func (self *Label) GetText() string {
	txt := C.gtk_label_get_text(self.object)
	return gobject.GoString(unsafe.Pointer(txt))
}

func (self *Label) SelectRegion(start_offset, end_offset int) {
	C.gtk_label_select_region(self.object, C.gint(start_offset), C.gint(end_offset))
}

func (self *Label) SetMnemonicWidget(w WidgetLike) {
	C.gtk_label_set_mnemonic_widget(self.object, w.W().object)
}

func (self *Label) SetSelectable(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_label_set_selectable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Label) SetTextWithMnemonic(text string) {
	s := gobject.GString(text)
	defer s.Free()
	C.gtk_label_set_text_with_mnemonic(self.object, (*C.gchar)(s.GetPtr()))
}

//TODO: gtk_label_get_attributes

func (self *Label) GetJustify() int {
	j := C.gtk_label_get_justify(self.object)
	return int(j)
}

//TODO: gtk_label_get_ellipsize

func (self *Label) GetWidthChars() int {
	w := C.gtk_label_get_width_chars(self.object)
	return int(w)
}

func (self *Label) GetMaxWidthChars() int {
	w := C.gtk_label_get_max_width_chars(self.object)
	return int(w)
}

func (self *Label) GetLabel() string {
	l := C.gtk_label_get_label(self.object)
	return gobject.GoString(unsafe.Pointer(&l))
}

//TODO: gtk_label_get_layout

func (self *Label) GetLineWrap() bool {
	lw := C.gtk_label_get_line_wrap(self.object)
	return gobject.GoBool(unsafe.Pointer(&lw))
}

//TODO: gtk_label_get_line_wrap_mode

func (self *Label) GetMnemonicWidget() WidgetLike {
	w := C.gtk_label_get_mnemonic_widget(self.object)
	if w == nil {
		return nil
	}

	widget, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err != nil {
		return widget.(WidgetLike)
	}

	return nil
}

func (self *Label) GetSelectionBounds() (bool, int, int) {
	var start C.gint
	var end C.gint
	b := C.gtk_label_get_selection_bounds(self.object, &start, &end)
	isSel := gobject.GoBool(unsafe.Pointer(&b))

	return isSel, int(start), int(end)
}

func (self *Label) GetUseMarkup() bool {
	b := C.gtk_label_get_use_markup(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Label) GetUseUnderline() bool {
	b := C.gtk_label_get_use_underline(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Label) GetSingleLineMode() bool {
	b := C.gtk_label_get_single_line_mode(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Label) GetAngle() float64 {
	e := C.gtk_label_get_angle(self.object)
	return float64(e)
}

func (self *Label) SetLabel(str string) {
	s := gobject.GString(str)
	defer s.Free()
	C.gtk_label_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Label) SetUseMarkup(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_label_set_use_markup(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Label) SetUseUnderline(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_label_set_use_underline(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Label) SetSingleLineMode(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_label_set_single_line_mode(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Label) SetAngle(angle float64) {
	C.gtk_label_set_angle(self.object, C.gdouble(angle))
}

func (self *Label) GetCurrentURI() string {
	uri := C.gtk_label_get_current_uri(self.object)
	return gobject.GoString(unsafe.Pointer(&uri))
}

func (self *Label) SetTrackVisitedLinks(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_label_set_track_visited_links(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Label) GetTrackVisitedLinks() bool {
	b := C.gtk_label_get_track_visited_links(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// END GtkLabel
////////////////////////////// }}}

// GtkProgressBar {{{
//////////////////////////////

// GtkProgressBar type
type ProgressBar struct {
	object *C.GtkProgressBar
	*Widget
	*Orientable
}

func NewProgressBar() *ProgressBar {
	pb := &ProgressBar{}
	o := C.gtk_progress_bar_new()
	pb.object = C.to_GtkProgressBar(unsafe.Pointer(o))

	if gobject.IsObjectFloating(pb) {
		gobject.RefSink(pb)
	}
	pb.Widget = NewWidget(unsafe.Pointer(o))
	pb.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	progressBarFinalizer(pb)

	return pb
}

// Clear ProgressBar struct when it goes out of reach
func progressBarFinalizer(pb *ProgressBar) {
	runtime.SetFinalizer(pb, func(pb *ProgressBar) { gobject.Unref(pb) })
}

// Conversion funcs
func newProgressBarFromNative(obj unsafe.Pointer) interface{} {
	pb := &ProgressBar{}
	pb.object = C.to_GtkProgressBar(obj)

	if gobject.IsObjectFloating(pb) {
		gobject.RefSink(pb)
	} else {
		gobject.Ref(pb)
	}
	pb.Widget = NewWidget(obj)
	pb.Orientable = newOrientableFromNative(obj).(*Orientable)
	progressBarFinalizer(pb)

	return pb
}

func nativeFromProgressBar(pb interface{}) *gobject.GValue {
	progress, ok := pb.(*ProgressBar)
	if ok {
		gv := gobject.CreateCGValue(GtkType.PROGRESS_BAR, progress.ToNative())
		return gv
	}
	return nil
}

// To be object like
func (self ProgressBar) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ProgressBar) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ProgressBar) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self ProgressBar) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self ProgressBar) W() *Widget {
	return self.Widget
}

// To be orientable-like
func (self ProgressBar) IOrientable() *Orientable {
	return self.Orientable
}

// ProgressBar interface

func (self *ProgressBar) Pulse() {
	C.gtk_progress_bar_pulse(self.object)
}

func (self *ProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(self.object, C.gdouble(fraction))
}

func (self *ProgressBar) GetFraction() float64 {
	return float64(C.gtk_progress_bar_get_fraction(self.object))
}

func (self *ProgressBar) SetInverted(inverted bool) {
	b := gobject.GBool(inverted)
	defer b.Free()
	C.gtk_progress_bar_set_inverted(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ProgressBar) GetInverted() bool {
	b := C.gtk_progress_bar_get_inverted(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ProgressBar) SetShowText(showText bool) {
	b := gobject.GBool(showText)
	defer b.Free()
	C.gtk_progress_bar_set_show_text(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ProgressBar) GetShowText() bool {
	b := C.gtk_progress_bar_get_show_text(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ProgressBar) SetText(text interface{}) {
	if text == nil {
		C.gtk_progress_bar_set_text(self.object, nil)
		return
	}
	if st, ok := text.(string); ok {
		s := gobject.GString(st)
		defer s.Free()
		C.gtk_progress_bar_set_text(self.object, (*C.gchar)(s.GetPtr()))
	}
}

func (self *ProgressBar) GetText() string {
	s := C.gtk_progress_bar_get_text(self.object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(s))
}

//TODO: gtk_progress_bar_set_ellipsize
//TODO: gtk_progress_bar_get_ellipsize

func (self *ProgressBar) SetPulseStep(fraction float64) {
	C.gtk_progress_bar_set_pulse_step(self.object, C.gdouble(fraction))
}

func (self *ProgressBar) GetPulseStep() float64 {
	return float64(C.gtk_progress_bar_get_pulse_step(self.object))
}

//////////////////////////////
// End GtkProgressBar
////////////////////////////// }}}

// GtkImage {{{
//////////////////////////////

// Image type
type Image struct {
	object *C.GtkImage
	*Widget
}

func NewImageFromStock(stockId string, iconSize int) *Image {
	im := &Image{}

	sId := gobject.GString(stockId)
	defer sId.Free()

	o := C.gtk_image_new_from_stock((*C.gchar)(sId.GetPtr()), C.GtkIconSize(iconSize))
	im.object = C.to_GtkImage(unsafe.Pointer(o))

	if gobject.IsObjectFloating(im) {
		gobject.RefSink(im)
	}
	im.Widget = NewWidget(unsafe.Pointer(o))
	imageFinalizer(im)

	return im
}

func NewImageFromFile(filename string) WidgetLike {
	im := &Image{}
	fn := gobject.GString(filename)
	defer fn.Free()

	o := C.gtk_image_new_from_file((*C.gchar)(fn.GetPtr()))
	im.object = C.to_GtkImage(unsafe.Pointer(o))

	if gobject.IsObjectFloating(im) {
		gobject.RefSink(im)
	}
	im.Widget = NewWidget(unsafe.Pointer(o))
	imageFinalizer(im)

	return im
}

// Clear Image struct when it goes out of reach
func imageFinalizer(im *Image) {
	runtime.SetFinalizer(im, func(im *Image) { gobject.Unref(im) })
}

// Conversion function for gobject registration map
func newImageFromNative(obj unsafe.Pointer) interface{} {
	im := &Image{}
	im.object = C.to_GtkImage(obj)

	if gobject.IsObjectFloating(im) {
		gobject.RefSink(im)
	} else {
		gobject.Ref(im)
	}
	im.Widget = NewWidget(obj)
	imageFinalizer(im)

	return im
}

func nativeFromImage(im interface{}) *gobject.GValue {
	image, ok := im.(*Image)
	if ok {
		gv := gobject.CreateCGValue(GtkType.IMAGE, image.ToNative())
		return gv
	}

	return nil
}

// To be object like
func (self Image) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Image) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Image) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self Image) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Image) W() *Widget {
	return self.Widget
}

//////////////////////////////
// END GtkImage
////////////////////////////// }}}

// GtkStatusbar {{{
//////////////////////////////

// GtkStatusbar type
type Statusbar struct {
	object *C.GtkStatusbar
	*Box
}

func NewStatusBar() *Statusbar {
	sb := &Statusbar{}
	o := C.gtk_statusbar_new()
	sb.object = C.to_GtkStatusbar(unsafe.Pointer(o))

	if gobject.IsObjectFloating(sb) {
		gobject.RefSink(sb)
	}
	sb.Box = newBoxFromNative(unsafe.Pointer(o)).(*Box)
	statusbarFinalizer(sb)

	return sb
}

// Clear Statusbar struct when it goes out of reach
func statusbarFinalizer(sb *Statusbar) {
	runtime.SetFinalizer(sb, func(sb *Statusbar) { gobject.Unref(sb) })
}

// Conversion funcs
func newStatusbarFromNative(obj unsafe.Pointer) interface{} {
	sb := &Statusbar{}
	sb.object = C.to_GtkStatusbar(obj)

	if gobject.IsObjectFloating(sb) {
		gobject.RefSink(sb)
	} else {
		gobject.Ref(sb)
	}
	sb.Box = newBoxFromNative(obj).(*Box)
	statusbarFinalizer(sb)

	return sb
}

func nativeFromStatusbar(sb interface{}) *gobject.GValue {
	status, ok := sb.(*Statusbar)
	if ok {
		gv := gobject.CreateCGValue(GtkType.STATUSBAR, status.ToNative())
		return gv
	}
	return nil
}

// To be object like
func (self Statusbar) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Statusbar) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Statusbar) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self Statusbar) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be box-like
func (self Statusbar) CBox() *Box {
	return self.Box
}

// Statusbar interface

func (self *Statusbar) GetContextId(contextDesc string) uint {
	s := gobject.GString(contextDesc)
	defer s.Free()
	return uint(C.gtk_statusbar_get_context_id(self.object, (*C.gchar)(s.GetPtr())))
}

func (self *Statusbar) Push(contextId uint, text string) uint {
	s := gobject.GString(text)
	defer s.Free()
	return uint(C.gtk_statusbar_push(self.object, C.guint(contextId), (*C.gchar)(s.GetPtr())))
}

func (self *Statusbar) Pop(contextId uint) {
	C.gtk_statusbar_pop(self.object, C.guint(contextId))
}

func (self *Statusbar) Remove(contextId, messageId uint) {
	C.gtk_statusbar_remove(self.object, C.guint(contextId), C.guint(messageId))
}

func (self *Statusbar) RemoveAll(contextId uint) {
	C.gtk_statusbar_remove_all(self.object, C.guint(contextId))
}

func (self *Statusbar) GetMessageArea() WidgetLike {
	w := C.gtk_statusbar_get_message_area(self.object)

	if w != nil {
		ma, err := gobject.ConvertToGo(unsafe.Pointer(w))
		if err == nil {
			return ma.(WidgetLike)
		}
	}
	return nil
}

//////////////////////////////
// End GtkStatusBar
////////////////////////////// }}}

// GtkSpinner {{{
//////////////////////////////

// Spinner type
type Spinner struct {
	object *C.GtkSpinner
	*Widget
}

func NewSpinner() *Spinner {
	s := &Spinner{}
	o := C.gtk_spinner_new()
	s.object = C.to_GtkSpinner(unsafe.Pointer(o))

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	}
	s.Widget = NewWidget(unsafe.Pointer(o))
	spinnerFinalizer(s)

	return s
}

// Clear Spinner struct when it goes out of reach
func spinnerFinalizer(s *Spinner) {
	runtime.SetFinalizer(s, func(s *Spinner) { gobject.Unref(s) })
}

// Conversion function for gobject registration map
func newSpinnerFromNative(obj unsafe.Pointer) interface{} {
	s := &Spinner{}
	s.object = C.to_GtkSpinner(obj)

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	} else {
		gobject.Ref(s)
	}
	s.Widget = NewWidget(obj)
	spinnerFinalizer(s)

	return s
}

func nativeFromSpinner(sp interface{}) *gobject.GValue {
	spinner, ok := sp.(*Spinner)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SPINNER, spinner.ToNative())
		return gv
	}
	return nil
}

// To be object like
func (self Spinner) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Spinner) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Spinner) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self Spinner) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Spinner) W() *Widget {
	return self.Widget
}

// Spinner interface

func (self *Spinner) Start() {
	C.gtk_spinner_start(self.object)
}

func (self *Spinner) Stop() {
	C.gtk_spinner_stop(self.object)
}

//////////////////////////////
// END GtkSpinner
////////////////////////////// }}}

// GtkInfoBar {{{
//////////////////////////////

// InfoBar type
type InfoBar struct {
	object *C.GtkInfoBar
	*Box
}

func NewInfoBar() *InfoBar {
    ib := &InfoBar{}
    o := C.gtk_info_bar_new()
    ib.object = C.to_GtkInfoBar(unsafe.Pointer(o))

    if gobject.IsObjectFloating(ib) {
        gobject.RefSink(ib)
    }
    ib.Box = newBoxFromNative(unsafe.Pointer(o)).(*Box)
    infoBarFinalizer(ib)

    return ib
}

func NewInfoBarWithButtons(buttons B) *InfoBar {
    ib := &InfoBar{}
    o := C.gtk_info_bar_new()
    ib.object = C.to_GtkInfoBar(unsafe.Pointer(o))

    if gobject.IsObjectFloating(ib) {
        gobject.RefSink(ib)
    }
    ib.Box = newBoxFromNative(unsafe.Pointer(o)).(*Box)
    infoBarFinalizer(ib)

    for _, butt := range buttons {
        ib.AddButton(butt.Text, butt.Response)
    }

    return ib
}

// Clear InfoBar struct when it goes out of reach
func infoBarFinalizer(ib *InfoBar) {
	runtime.SetFinalizer(ib, func(ib *InfoBar) { gobject.Unref(ib) })
}

// Conversion function for gobject registration map
func newInfoBarFromNative(obj unsafe.Pointer) interface{} {
	ib := &InfoBar{}
	ib.object = C.to_GtkInfoBar(obj)

	if gobject.IsObjectFloating(ib) {
		gobject.RefSink(ib)
	} else {
		gobject.Ref(ib)
	}
    ib.Box = newBoxFromNative(obj).(*Box)
    infoBarFinalizer(ib)

	return ib
}

func nativeFromInfoBar(ib interface{}) *gobject.GValue {
	info, ok := ib.(*InfoBar)
	if ok {
		gv := gobject.CreateCGValue(GtkType.INFO_BAR, info.ToNative())
		return gv
	}
	return nil
}

// To be object like
func (self InfoBar) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self InfoBar) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self InfoBar) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self InfoBar) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be box-like
func (self InfoBar) B() *Box {
    return self.Box
}

// InfoBar interface

func (self *InfoBar) AddButton(text string, responseId int) WidgetLike {
    s := gobject.GString(text)
    defer s.Free()
    w := C.gtk_info_bar_add_button(self.object, (*C.gchar)(s.GetPtr()), C.gint(responseId))

    if w != nil {
        if button, err := gobject.ConvertToGo(unsafe.Pointer(w)); err != nil {
            return button.(WidgetLike)
        }
    }

    return nil
}

func (self *InfoBar) AddButtons(buttons B) {
    for _, butt := range buttons {
        self.AddButton(butt.Text, butt.Response)
    }
}

func (self *InfoBar) AddActionWidget(child WidgetLike, responseId int) {
    C.gtk_info_bar_add_action_widget(self.object, child.W().object, C.gint(responseId))
}

func (self *InfoBar) SetResponseSensitive(responseId int, setting bool) {
    b := gobject.GBool(setting)
    defer b.Free()
    C.gtk_info_bar_set_response_sensitive(self.object, C.gint(responseId), *((*C.gboolean)(b.GetPtr())))
}

func (self *InfoBar) SetDefaultResponse(responseId int) {
    C.gtk_info_bar_set_default_response(self.object, C.gint(responseId))
}

func (self *InfoBar) Response(responseId int) {
    C.gtk_info_bar_response(self.object, C.gint(responseId))
}

func (self *InfoBar) SetMessageType(gtk_MessageType int) {
    C.gtk_info_bar_set_message_type(self.object, C.GtkMessageType(gtk_MessageType))
}

func (self *InfoBar) GetMessageType() int {
    i := C.gtk_info_bar_get_message_type(self.object)
    return int(i)
}

func (self *InfoBar) GetActionArea() WidgetLike {
    a := C.gtk_info_bar_get_action_area(self.object)

    if a != nil {
        if area, err := gobject.ConvertToGo(unsafe.Pointer(a)); err == nil {
            return area.(WidgetLike)
        }
    }

    return nil
}

func (self *InfoBar) GetContentArea() WidgetLike {
    c := C.gtk_info_bar_get_content_area(self.object)

    if c != nil {
        if area, err := gobject.ConvertToGo(unsafe.Pointer(c)); err == nil {
            return area.(WidgetLike)
        }
    }

    return nil
}

//////////////////////////////
// END GtkInfoBar
////////////////////////////// }}}

// End Display Widgets }}}

// Buttons and Toggles {{{

// GtkButton {{{
//////////////////////////////

// GtkButton type
type Button struct {
	object *C.GtkButton
	*Bin
}

//Create and return new button Structure
func NewButton() *Button {
	b := &Button{}
	o := C.gtk_button_new()
	b.object = C.to_GtkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Bin = NewBin(unsafe.Pointer(o))
	buttonFinalizer(b)

	return b
}

func NewButtonWithLabel(label string) *Button {
	b := &Button{}
	l := gobject.GString(label)
	defer l.Free()

	o := C.gtk_button_new_with_label((*C.gchar)(l.GetPtr()))
	b.object = C.to_GtkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Bin = NewBin(unsafe.Pointer(o))
	buttonFinalizer(b)

	return b
}

func NewButtonFromStock(stockId string) *Button {
	b := NewButtonWithLabel(stockId)
	b.SetUseStock(true)
	return b
}

func NewButtonWithMnemonic(label string) *Button {
	b := &Button{}
	l := gobject.GString(label)
	defer l.Free()
	o := C.gtk_button_new_with_mnemonic((*C.gchar)(l.GetPtr()))
	b.object = C.to_GtkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Bin = NewBin(unsafe.Pointer(o))
	buttonFinalizer(b)

	return b
}

// Unref button, when Button structer goes out of reach
func buttonFinalizer(b *Button) {
	runtime.SetFinalizer(b, func(b *Button) { gobject.Unref(b) })
}

// Conversion function for gobject registration map
func newButtonFromNative(obj unsafe.Pointer) interface{} {
	b := &Button{}
	b.object = C.to_GtkButton(obj)

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	} else {
		gobject.Ref(b)
	}
	b.Bin = NewBin(obj)
	buttonFinalizer(b)

	return b
}

func nativeFromButton(b interface{}) *gobject.GValue {
	if but, ok := b.(*Button); ok {
		gv := gobject.CreateCGValue(GtkType.BUTTON, but.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Button) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Button) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Button) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Button) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container like
func (self Button) CBin() *Bin {
	return self.Bin
}

// Button interface
func (self *Button) SetLabel(label string) {
	s := gobject.GString(label)
	defer s.Free()
	C.gtk_button_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Button) GetLabel() string {
	l := C.gtk_button_get_label(self.object)
	return gobject.GoString(unsafe.Pointer(l))
}

func (self *Button) GetUseStock() bool {
	b := C.gtk_button_get_use_stock(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Button) SetUseStock(useStock bool) {
	gobject.SetProperty(self, "use-stock", useStock)
}

func (self *Button) GetUseUnderline() bool {
	b := C.gtk_button_get_use_underline(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Button) SetUseUnderline(underline bool) {
	gobject.SetProperty(self, "use-underline", underline)
}

func (self *Button) SetFocusOnClick(focusOnClick bool) {
	gobject.SetProperty(self, "focus-on-click", focusOnClick)
}

func (self *Button) GetFocusOnClick() bool {
	b := C.gtk_button_get_focus_on_click(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Button) SetAlignment(xalign, yalign float32) {
	C.gtk_button_set_alignment(self.object, C.gfloat(xalign), C.gfloat(yalign))
}

func (self *Button) GetAlignment() (float32, float32) {
	var cx C.gfloat
	var cy C.gfloat
	C.gtk_button_get_alignment(self.object, &cx, &cy)
	return float32(cx), float32(cy)
}

func (self *Button) SetImage(w WidgetLike) {
	C.gtk_button_set_image(self.object, w.W().object)
}

func (self *Button) GetImage() WidgetLike {
	cw := C.gtk_button_get_image(self.object)
	widget, err := gobject.ConvertToGo(unsafe.Pointer(&cw))
	if err != nil {
		return widget.(WidgetLike)
	}

	return nil
}

func (self *Button) SetImagePosition(position int) {
	C.gtk_button_set_image_position(self.object, C.GtkPositionType(position))
}

func (self *Button) GetImagePosition() int {
	return int(C.gtk_button_get_image_position(self.object))
}

func (self *Button) Clicked() {
	C.gtk_button_clicked(self.object)
}

//TODO: gtk_button_get_event_window
//////////////////////////////
// END GtkButton
////////////////////////////// }}}

// GtkToggleButton {{{
//////////////////////////////

// GtkToggleButton type
type ToggleButton struct {
	object *C.GtkToggleButton
	*Button
}

// Create and return new toggle button
func NewToggleButton() *ToggleButton {
	b := &ToggleButton{}
	o := C.gtk_toggle_button_new()
	b.object = C.to_GtkToggleButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	toggleButtonFinalizer(b)

	return b
}

func NewToggleButtonWithLabel(label string) *ToggleButton {
	b := &ToggleButton{}
	l := gobject.GString(label)
	defer l.Free()

	o := C.gtk_toggle_button_new_with_label((*C.gchar)(l.GetPtr()))
	b.object = C.to_GtkToggleButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	toggleButtonFinalizer(b)

	return b
}

// Clear ToggleButton when it goes out of reach
func toggleButtonFinalizer(b *ToggleButton) {
	runtime.SetFinalizer(b, func(b *ToggleButton) { gobject.Unref(b) })
}

// Conversion function for gobject registration map
func newToggleButtonFromNative(obj unsafe.Pointer) interface{} {
	b := &ToggleButton{}
	b.object = C.to_GtkToggleButton(obj)

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	} else {
		gobject.Ref(b)
	}
	b.Button = newButtonFromNative(obj).(*Button)
	toggleButtonFinalizer(b)

	return b
}

func nativeFromToggleButton(b interface{}) *gobject.GValue {
	if but, ok := b.(*ToggleButton); ok {
		gv := gobject.CreateCGValue(GtkType.TOGGLE_BUTTON, but.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self ToggleButton) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ToggleButton) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ToggleButton) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ToggleButton) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be button-like
func (self ToggleButton) B() *Button {
	return self.Button
}

// ToggleButton interface

func (self *ToggleButton) SetMode(drawIndicator bool) {
	b := gobject.GBool(drawIndicator)
	defer b.Free()
	C.gtk_toggle_button_set_mode(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ToggleButton) GetMode() bool {
	cb := C.gtk_toggle_button_get_mode(self.object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ToggleButton) Toggled() {
	C.gtk_toggle_button_toggled(self.object)
}

func (self *ToggleButton) GetActive() bool {
	cb := C.gtk_toggle_button_get_active(self.object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ToggleButton) SetActive(isActive bool) {
	b := gobject.GBool(isActive)
	defer b.Free()
	C.gtk_toggle_button_set_active(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ToggleButton) GetInconsistent() bool {
	cb := C.gtk_toggle_button_get_inconsistent(self.object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ToggleButton) SetInconsistent(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_toggle_button_set_inconsistent(self.object, *((*C.gboolean)(b.GetPtr())))
}

//////////////////////////////
// END GtkToggleButton
////////////////////////////// }}}

// GtkCheckButton {{{
//////////////////////////////

// GtkCheckButton type
type CheckButton struct {
	object *C.GtkCheckButton
	*ToggleButton
}

// Clear CheckButton struct when it goes out of reach
func checkButtonFinalizer(cb *CheckButton) {
	runtime.SetFinalizer(cb, func(cb *CheckButton) { gobject.Unref(cb) })
}

func NewCheckButton() *CheckButton {
	cb := &CheckButton{}
	o := C.gtk_check_button_new()
	cb.object = C.to_GtkCheckButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.ToggleButton = newToggleButtonFromNative(unsafe.Pointer(o)).(*ToggleButton)
	checkButtonFinalizer(cb)

	return cb
}

func NewCheckButtonWithLabel(label string) *CheckButton {
	cb := &CheckButton{}
	s := gobject.GString(label)
	defer s.Free()

	o := C.gtk_check_button_new_with_label((*C.gchar)(s.GetPtr()))
	cb.object = C.to_GtkCheckButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.ToggleButton = newToggleButtonFromNative(unsafe.Pointer(o)).(*ToggleButton)
	checkButtonFinalizer(cb)

	return cb
}

func NewCheckButtonWithMnemonic(label string) *CheckButton {
	cb := &CheckButton{}
	s := gobject.GString(label)
	defer s.Free()

	o := C.gtk_check_button_new_with_mnemonic((*C.gchar)(s.GetPtr()))
	cb.object = C.to_GtkCheckButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.ToggleButton = newToggleButtonFromNative(unsafe.Pointer(o)).(*ToggleButton)
	checkButtonFinalizer(cb)

	return cb
}

// Conversion funcs
func newCheckButtonFromNative(obj unsafe.Pointer) interface{} {
	cb := &CheckButton{}
	cb.object = C.to_GtkCheckButton(obj)

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	} else {
		gobject.Ref(cb)
	}
	cb.ToggleButton = newToggleButtonFromNative(obj).(*ToggleButton)
	checkButtonFinalizer(cb)

	return cb
}

func nativeFromCheckButton(cb interface{}) *gobject.GValue {
	check, ok := cb.(*CheckButton)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CHECK_BUTTON, check.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CheckButton) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CheckButton) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CheckButton) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CheckButton) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

//////////////////////////////
// END GtkCheckButton
////////////////////////////// }}}

// GtkRadioButton {{{
//////////////////////////////

// GtkRadioButton type
type RadioButton struct {
	object *C.GtkRadioButton
	*CheckButton
}

// Clear RadioButton struct when it goes out of reach
func radioButtonFinalizer(rb *RadioButton) {
	runtime.SetFinalizer(rb, func(rb *RadioButton) { gobject.Unref(rb) })
}

func NewRadioButton(group *glib.GSList) *RadioButton {
	var list *C.GSList = nil
	if group != nil {
		list = (*C.GSList)(group.ToNative())
	}
	rb := &RadioButton{}
	o := C.gtk_radio_button_new(list)
	rb.object = C.to_GtkRadioButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(rb) {
		gobject.RefSink(rb)
	}
	rb.CheckButton = newCheckButtonFromNative(unsafe.Pointer(o)).(*CheckButton)
	radioButtonFinalizer(rb)

	return rb
}

func NewRadioButtonFromWidget(radioGroupMember *RadioButton) *RadioButton {
	if radioGroupMember == nil {
		return NewRadioButton(nil)
	}

	rb := &RadioButton{}
	o := C.gtk_radio_button_new_from_widget(radioGroupMember.object)
	rb.object = C.to_GtkRadioButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(rb) {
		gobject.RefSink(rb)
	}
	rb.CheckButton = newCheckButtonFromNative(unsafe.Pointer(o)).(*CheckButton)
	radioButtonFinalizer(rb)

	return rb
}

func NewRadioButtonWithLabel(group *glib.GSList, label string) *RadioButton {
	var list *C.GSList = nil
	if group != nil {
		list = (*C.GSList)(group.ToNative())
	}
	s := gobject.GString(label)
	defer s.Free()

	rb := &RadioButton{}

	o := C.gtk_radio_button_new_with_label(list, (*C.gchar)(s.GetPtr()))
	rb.object = C.to_GtkRadioButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(rb) {
		gobject.RefSink(rb)
	}
	rb.CheckButton = newCheckButtonFromNative(unsafe.Pointer(o)).(*CheckButton)
	radioButtonFinalizer(rb)

	return rb
}

func NewRadioButtonWithLabelFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	if radioGroupMember == nil {
		return NewRadioButtonWithLabel(nil, label)
	}

	s := gobject.GString(label)
	defer s.Free()

	rb := &RadioButton{}
	o := C.gtk_radio_button_new_with_label_from_widget(radioGroupMember.object, (*C.gchar)(s.GetPtr()))
	rb.object = C.to_GtkRadioButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(rb) {
		gobject.RefSink(rb)
	}
	rb.CheckButton = newCheckButtonFromNative(unsafe.Pointer(o)).(*CheckButton)
	radioButtonFinalizer(rb)

	return rb
}

func NewRadioButtonWithMnemonic(group *glib.GSList, label string) *RadioButton {
	var list *C.GSList = nil
	if group != nil {
		list = (*C.GSList)(group.ToNative())
	}
	s := gobject.GString(label)
	defer s.Free()

	rb := &RadioButton{}

	o := C.gtk_radio_button_new_with_mnemonic(list, (*C.gchar)(s.GetPtr()))
	rb.object = C.to_GtkRadioButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(rb) {
		gobject.RefSink(rb)
	}
	rb.CheckButton = newCheckButtonFromNative(unsafe.Pointer(o)).(*CheckButton)
	radioButtonFinalizer(rb)

	return rb
}

func NewRadioButtonWithMnemonicFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	if radioGroupMember == nil {
		return NewRadioButtonWithMnemonic(nil, label)
	}

	s := gobject.GString(label)
	defer s.Free()

	rb := &RadioButton{}
	o := C.gtk_radio_button_new_with_mnemonic_from_widget(radioGroupMember.object, (*C.gchar)(s.GetPtr()))
	rb.object = C.to_GtkRadioButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(rb) {
		gobject.RefSink(rb)
	}
	rb.CheckButton = newCheckButtonFromNative(unsafe.Pointer(o)).(*CheckButton)
	radioButtonFinalizer(rb)

	return rb
}

// Conversion funcs
func newRadioButtonFromNative(obj unsafe.Pointer) interface{} {
	rb := &RadioButton{}
	rb.object = C.to_GtkRadioButton(obj)

	if gobject.IsObjectFloating(rb) {
		gobject.RefSink(rb)
	} else {
		gobject.Ref(rb)
	}
	rb.CheckButton = newCheckButtonFromNative(obj).(*CheckButton)
	radioButtonFinalizer(rb)

	return rb
}

func nativeFromRadioButton(rb interface{}) *gobject.GValue {
	radio, ok := rb.(*RadioButton)
	if ok {
		gv := gobject.CreateCGValue(GtkType.RADIO_BUTTON, radio.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self RadioButton) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self RadioButton) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self RadioButton) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self RadioButton) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// RadioButton interface

func (self *RadioButton) SetGroup(group *glib.GSList) {
	C.gtk_radio_button_set_group(self.object, (*C.GSList)(group.ToNative()))
}

func (self *RadioButton) GetGroup() *glib.GSList {
	g := C.gtk_radio_button_get_group(self.object)

	if g != nil {
		goList := glib.NewGSListFromNative(unsafe.Pointer(g))
		goList.GC_Free = false
		goList.GC_FreeFull = false
		goList.ConversionFunc = newRadioButtonFromNative
		return goList
	}
	return nil
}

func (self *RadioButton) JoinGroup(groupSource *RadioButton) {
	if groupSource == nil {
		C.gtk_radio_button_join_group(self.object, nil)
	}
	C.gtk_radio_button_join_group(self.object, groupSource.object)
}

//////////////////////////////
// END GtkRadioButton
////////////////////////////// }}}

// GtkLinkButton {{{
//////////////////////////////

// GtkLinkButton type
type LinkButton struct {
	object *C.GtkLinkButton
	*Button
}

// Create and return new link button
func NewLinkButton(uri string) *LinkButton {
	s := gobject.GString(uri)
	defer s.Free()

	b := &LinkButton{}
	o := C.gtk_link_button_new((*C.gchar)(s.GetPtr()))
	b.object = C.to_GtkLinkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	linkButtonFinalizer(b)

	return b
}

func NewLinkButtonWithLabel(uri, label string) *LinkButton {
	b := &LinkButton{}

	u := gobject.GString(uri)
	defer u.Free()

	l := gobject.GString(label)
	defer l.Free()

	o := C.gtk_link_button_new_with_label((*C.gchar)(u.GetPtr()), (*C.gchar)(l.GetPtr()))
	b.object = C.to_GtkLinkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	linkButtonFinalizer(b)

	return b
}

// Clear LinkButton when it goes out of reach
func linkButtonFinalizer(b *LinkButton) {
	runtime.SetFinalizer(b, func(b *LinkButton) { gobject.Unref(b) })
}

// Conversion function for gobject registration map
func newLinkButtonFromNative(obj unsafe.Pointer) interface{} {
	b := &LinkButton{}
	b.object = C.to_GtkLinkButton(obj)

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	} else {
		gobject.Ref(b)
	}
	b.Button = newButtonFromNative(obj).(*Button)
	linkButtonFinalizer(b)

	return b
}

func nativeFromLinkButton(b interface{}) *gobject.GValue {
	if but, ok := b.(*LinkButton); ok {
		gv := gobject.CreateCGValue(GtkType.LINK_BUTTON, but.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self LinkButton) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self LinkButton) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self LinkButton) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self LinkButton) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be button-like
func (self LinkButton) B() *Button {
	return self.Button
}

// LinkButton interface

func (self *LinkButton) GetURI() string {
	s := C.gtk_link_button_get_uri(self.object)

	if s != nil {
		return gobject.GoString(unsafe.Pointer(C.g_strdup(s)))
	}
	return ""
}

func (self *LinkButton) SetURI(uri string) {
	s := gobject.GString(uri)
	defer s.Free()
	C.gtk_link_button_set_uri(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *LinkButton) GetVisited() bool {
	b := C.gtk_link_button_get_visited(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *LinkButton) SetVisited(visited bool) {
	b := gobject.GBool(visited)
	defer b.Free()
	C.gtk_link_button_set_visited(self.object, *((*C.gboolean)(b.GetPtr())))
}

//////////////////////////////
// END GtkLinkButton
////////////////////////////// }}}

// GtkSwitch {{{
//////////////////////////////

// GtkSwitch type
type Switch struct {
	object *C.GtkSwitch
	*Widget
}

func NewSwitch() *Switch {
	s := &Switch{}

	o := C.gtk_switch_new()
	s.object = C.to_GtkSwitch(unsafe.Pointer(o))

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	}
	s.Widget = NewWidget(unsafe.Pointer(o))
	switchFinalizer(s)

	return s
}

// Clear switch struct when it goes out of reach
func switchFinalizer(s *Switch) {
	runtime.SetFinalizer(s, func(s *Switch) { gobject.Unref(s) })
}

// Conversion function for gobject registration map
func newSwitchFromNative(obj unsafe.Pointer) interface{} {
	s := &Switch{}
	s.object = C.to_GtkSwitch(obj)

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	} else {
		gobject.Ref(s)
	}
	s.Widget = NewWidget(unsafe.Pointer(s.object))
	switchFinalizer(s)

	return s
}

func nativeFromSwitch(sw interface{}) *gobject.GValue {
	s, ok := sw.(*Switch)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SWITCH, s.ToNative())
		return gv
	}

	return nil
}

// To be object like
func (self Switch) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Switch) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Switch) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self Switch) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Switch) W() *Widget {
	return self.Widget
}

// Switch interface

func (self *Switch) SetActive(isActive bool) {
    b := gobject.GBool(isActive)
    defer b.Free()
    C.gtk_switch_set_active(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Switch) GetActive() bool {
    b := C.gtk_switch_get_active(self.object)
    return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// END GtkSwitch
////////////////////////////// }}}

// End Buttons and Toggles }}}

// Numeric/Text Data Entry {{{

// GtkEntryBuffer {{{
//////////////////////////////

// GtkEntryBuffer type
type EntryBuffer struct {
	object *C.GtkEntryBuffer
}

func NewEntryBuffer(initial_chars string) *EntryBuffer {
	e := &EntryBuffer{}

	s := gobject.GString(initial_chars)
	defer s.Free()

	o := C.gtk_entry_buffer_new((*C.gchar)(s.GetPtr()), C.gint(len(initial_chars)))
	e.object = C.to_GtkEntryBuffer(unsafe.Pointer(o))

	if gobject.IsObjectFloating(e) {
		gobject.RefSink(e)
	}
	entryBufferFinalizer(e)

	return e

}

// Clear EntryBuffer struct when it goes out of reach
func entryBufferFinalizer(e *EntryBuffer) {
	runtime.SetFinalizer(e, func(e *EntryBuffer) { gobject.Unref(e) })
}

// Conversion function for gobject registration map
func newEntryBufferFromNative(obj unsafe.Pointer) interface{} {
	e := &EntryBuffer{}
	e.object = C.to_GtkEntryBuffer(obj)

	if gobject.IsObjectFloating(e) {
		gobject.RefSink(e)
	} else {
		gobject.Ref(e)
	}

	return e
}

func nativeFromEntryBuffer(eb interface{}) *gobject.GValue {
	e, ok := eb.(*EntryBuffer)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ENTRY_BUFFER, e.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self EntryBuffer) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self EntryBuffer) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self EntryBuffer) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self EntryBuffer) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// EntryBuffer interface

func (self *EntryBuffer) GetText() string {
	s := C.gtk_entry_buffer_get_text(self.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *EntryBuffer) SetText(chars string) {
	s := gobject.GString(chars)
	defer s.Free()

	C.gtk_entry_buffer_set_text(self.object, (*C.gchar)(s.GetPtr()), C.gint(len(chars)))
}

func (self *EntryBuffer) GetBytes() uint64 {
	return uint64(C.gtk_entry_buffer_get_bytes(self.object))
}

func (self *EntryBuffer) GetLength() uint {
	return uint(C.gtk_entry_buffer_get_length(self.object))
}

func (self *EntryBuffer) GetMaxLength() int {
	return int(C.gtk_entry_buffer_get_max_length(self.object))
}

func (self *EntryBuffer) SetMaxLength(maxLength int) {
	C.gtk_entry_buffer_set_max_length(self.object, C.gint(maxLength))
}

func (self *EntryBuffer) InsertText(position uint, chars string) {
	s := gobject.GString(chars)
	defer s.Free()

	C.gtk_entry_buffer_insert_text(self.object, C.guint(position), (*C.gchar)(s.GetPtr()), C.gint(len(chars)))
}

func (self *EntryBuffer) DeleteText(position uint, num_chars int) {
	C.gtk_entry_buffer_delete_text(self.object, C.guint(position), C.gint(num_chars))
}

//////////////////////////////
// END GtkEntryBuffer
////////////////////////////// }}}

// GtkEntry {{{
//////////////////////////////

// GtkEntry Type
type Entry struct {
	object *C.GtkEntry
	*Widget
}

func NewEntry() *Entry {
	e := &Entry{}
	o := C.gtk_entry_new()
	e.object = C.to_GtkEntry(unsafe.Pointer(o))

	if gobject.IsObjectFloating(e) {
		gobject.RefSink(e)
	}
	e.Widget = NewWidget(unsafe.Pointer(o))
	entryFinalizer(e)

	return e
}

func NewEntryWithBuffer(eb *EntryBuffer) *Entry {
	e := &Entry{}
	o := C.gtk_entry_new_with_buffer(eb.object)
	e.object = C.to_GtkEntry(unsafe.Pointer(o))

	if gobject.IsObjectFloating(e) {
		gobject.RefSink(e)
	}
	e.Widget = NewWidget(unsafe.Pointer(o))
	entryFinalizer(e)

	return e
}

// Clear Entry struct when it goes out of reach
func entryFinalizer(e *Entry) {
	runtime.SetFinalizer(e, func(e *Entry) { gobject.Unref(e) })
}

// Conversion function for gobject registration map
func newEntryFromNative(obj unsafe.Pointer) interface{} {
	e := &Entry{}
	e.object = C.to_GtkEntry(obj)

	if gobject.IsObjectFloating(e) {
		gobject.RefSink(e)
	} else {
		gobject.Ref(e)
	}
	e.Widget = NewWidget(obj)
	entryFinalizer(e)

	return e
}

func nativeFromEntry(entry interface{}) *gobject.GValue {
	e, ok := entry.(*Entry)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ENTRY, e.ToNative())
		return gv
	}
	return nil
}

// To be Object-like
func (self Entry) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Entry) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Entry) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Entry) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Entry) W() *Widget {
	return self.Widget
}

// Entry interface
func (self *Entry) GetBuffer() *EntryBuffer {
	b := C.gtk_entry_get_buffer(self.object)
	buf, err := gobject.ConvertToGo(unsafe.Pointer(b))
	if err == nil {
		return buf.(*EntryBuffer)
	}
	return nil
}

func (self *Entry) SetBuffer(buffer *EntryBuffer) {
	C.gtk_entry_set_buffer(self.object, buffer.object)
}

func (self *Entry) SetText(text string) {
	s := gobject.GString(text)
	defer s.Free()

	C.gtk_entry_set_text(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Entry) GetText() string {
	s := C.gtk_entry_get_text(self.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Entry) GetTextLength() uint {
	return uint(C.gtk_entry_get_text_length(self.object))
}

func (self *Entry) GetTextArea() gdk3.Rectangle {
	var crec C.GdkRectangle
	C.gtk_entry_get_text_area(self.object, &crec)

	rec, err := gobject.ConvertToGo(unsafe.Pointer(&crec), gdk3.GdkType.RECTANGLE)
	if err == nil {
		return rec.(gdk3.Rectangle)
	}
	return gdk3.Rectangle{}
}

func (self *Entry) SetVisibility(visible bool) {
	b := gobject.GBool(visible)
	defer b.Free()
	C.gtk_entry_set_visibility(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Entry) SetInvisiblechar(char rune) {
	C.gtk_entry_set_invisible_char(self.object, C.gunichar(char))
}

func (self *Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(self.object)
}

func (self *Entry) SetMaxLength(max int) {
	C.gtk_entry_set_max_length(self.object, C.gint(max))
}

func (self *Entry) GetActivatesDefault() bool {
	b := C.gtk_entry_get_activates_default(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Entry) GetHasFrame() bool {
	b := C.gtk_entry_get_has_frame(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//TODO: gtk_entry_get_inner_border

func (self *Entry) GetWidthChars() int {
	return int(C.gtk_entry_get_width_chars(self.object))
}

func (self *Entry) SetActivatesDefault(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_entry_set_activates_default(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Entry) SetHasFrame(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_entry_set_has_frame(self.object, *((*C.gboolean)(b.GetPtr())))
}

//TODO: gtk_entry_set_inner_border

func (self *Entry) SetWidthChars(nChars int) {
	C.gtk_entry_set_width_chars(self.object, C.gint(nChars))
}

func (self *Entry) GetInvisibleChar() rune {
	return rune(C.gtk_entry_get_invisible_char(self.object))
}

func (self *Entry) SetAlignment(xalign float32) {
	C.gtk_entry_set_alignment(self.object, C.gfloat(xalign))
}

func (self *Entry) GetAlignment() float32 {
	return float32(C.gtk_entry_get_alignment(self.object))
}

func (self *Entry) SetPlaceholderText(text string) {
	s := gobject.GString(text)
	defer s.Free()
	C._gtk_entry_set_placeholder_text(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Entry) GetPlaceholderText() string {
	s := C._gtk_entry_get_placeholder_text(self.object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Entry) SetOverwriteMode(overwrite bool) {
	b := gobject.GBool(overwrite)
	defer b.Free()
	C.gtk_entry_set_overwrite_mode(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Entry) GetOverwriteMode() bool {
	b := C.gtk_entry_get_overwrite_mode(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//TODO: gtk_entry_get_layout

func (self *Entry) GetLayoutOffsets() (x, y int) {
	var cx, cy C.gint
	C.gtk_entry_get_layout_offsets(self.object, &cx, &cy)
	x = int(cx)
	y = int(cy)
	return
}

func (self *Entry) LayoutIndexToTextIndex(layoutIndex int) {
	C.gtk_entry_layout_index_to_text_index(self.object, C.gint(layoutIndex))
}

func (self *Entry) TextIndexToLayoutIndex(textIndex int) {
	C.gtk_entry_text_index_to_layout_index(self.object, C.gint(textIndex))
}

func (self *Entry) GetMaxLength() int {
	return int(C.gtk_entry_get_max_length(self.object))
}

func (self *Entry) GetVisibility() bool {
	b := C.gtk_entry_get_visibility(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Entry) SetCompletion(completion *EntryCompletion) {
	C.gtk_entry_set_completion(self.object, completion.object)
}

func (self *Entry) GetCompletion() *EntryCompletion {
	c := C.gtk_entry_get_completion(self.object)

	if c != nil {
		if com, err := gobject.ConvertToGo(unsafe.Pointer(c)); err == nil {
			return com.(*EntryCompletion)
		}
	}
	return nil
}

func (self *Entry) SetCursorHadjustment(adj *Adjustment) {
	if adj == nil {
		C.gtk_entry_set_cursor_hadjustment(self.object, nil)
		return
	}
	C.gtk_entry_set_cursor_hadjustment(self.object, adj.object)
}

func (self *Entry) GetCursorHadjustment() *Adjustment {
	h := C.gtk_entry_get_cursor_hadjustment(self.object)
	if h != nil {
		if adj, err := gobject.ConvertToGo(unsafe.Pointer(h)); err == nil {
			return adj.(*Adjustment)
		}
	}
	return nil
}

func (self *Entry) SetProgressFraction(fraction float64) {
	C.gtk_entry_set_progress_fraction(self.object, C.gdouble(fraction))
}

func (self *Entry) GetProgressFraction() float64 {
	return float64(C.gtk_entry_get_progress_fraction(self.object))
}

func (self *Entry) SetProgressPulseStep(fraction float64) {
	C.gtk_entry_set_progress_pulse_step(self.object, C.gdouble(fraction))
}

func (self *Entry) GetProgressPulseStep() float64 {
	return float64(C.gtk_entry_get_progress_pulse_step(self.object))
}

func (self *Entry) ProgressPulse() {
	C.gtk_entry_progress_pulse(self.object)
}

//TODO: gkt_entry_im_context_filter_keypress

func (self *Entry) ResetImContext() {
	C.gtk_entry_reset_im_context(self.object)
}

func (self *Entry) SetIconFromPixbuf(gtk_EntryIconPosition int, pixbuf *gdkpixbuf.Pixbuf) {
	if pixbuf == nil {
		C.gtk_entry_set_icon_from_pixbuf(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition), nil)
		return
	}
	C.gtk_entry_set_icon_from_pixbuf(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition),
		(*C.GdkPixbuf)(pixbuf.ToNative()))
}

func (self *Entry) SetIconFromStock(gtk_EntryIconPosition int, stockId interface{}) {
	if stockId == nil {
		C.gtk_entry_set_icon_from_stock(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition), nil)
		return
	}

	if sId, ok := stockId.(string); ok {
		s := gobject.GString(sId)
		defer s.Free()
		C.gtk_entry_set_icon_from_stock(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition),
			(*C.gchar)(s.GetPtr()))
	}
}

func (self *Entry) SetIconFromIconName(gtk_EntryIconPosition int, iconName interface{}) {
	if iconName == nil {
		C.gtk_entry_set_icon_from_icon_name(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition), nil)
		return
	}
	if name, ok := iconName.(string); ok {
		s := gobject.GString(name)
		defer s.Free()
		C.gtk_entry_set_icon_from_stock(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition),
			(*C.gchar)(s.GetPtr()))
	}
}

//TODO: gtk_entry_set_icon_from_gicon

func (self *Entry) GetIconStorageType(gtk_EntryIconPosition int) int {
	return int(C.gtk_entry_get_icon_storage_type(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition)))
}

func (self *Entry) GetIconPixbuf(gtk_EntryIconPosition int) *gdkpixbuf.Pixbuf {
	p := C.gtk_entry_get_icon_pixbuf(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition))

	if p != nil {
		if pix, err := gobject.ConvertToGo(unsafe.Pointer(p)); err == nil {
			return pix.(*gdkpixbuf.Pixbuf)
		}
	}
	return nil
}

func (self *Entry) GetIconStock(gtk_EntryIconPosition int) string {
	s := C.gtk_entry_get_icon_stock(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition))

	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}

func (self *Entry) GetIconName(gtk_EntryIconPosition int) string {
	s := C.gtk_entry_get_icon_name(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition))

	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}

//TODO: gtk_entry_get_icon_gicon

func (self *Entry) SetIconActivatable(gtk_EntryIconPosition int, activatable bool) {
	b := gobject.GBool(activatable)
	defer b.Free()
	C.gtk_entry_set_icon_activatable(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition),
		*((*C.gboolean)(b.GetPtr())))
}

func (self *Entry) GetIconActivatable(gtk_EntryIconPosition int) bool {
	b := C.gtk_entry_get_icon_activatable(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Entry) SetIconSensitive(gtk_EntryIconPosition int, sensitive bool) {
	b := gobject.GBool(sensitive)
	defer b.Free()
	C.gtk_entry_set_icon_sensitive(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition),
		*((*C.gboolean)(b.GetPtr())))
}

func (self *Entry) GetIconSensitive(gtk_EntryIconPosition int) bool {
	b := C.gtk_entry_get_icon_sensitive(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Entry) GetIconAtPos(x, y int) int {
	return int(C.gtk_entry_get_icon_at_pos(self.object, C.gint(x), C.gint(y)))
}

func (self *Entry) SetIconTooltipText(gtk_EntryIconPosition int, tooltip interface{}) {
	if tooltip == nil {
		C.gtk_entry_set_icon_tooltip_text(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition), nil)
		return
	}

	if tt, ok := tooltip.(string); ok {
		s := gobject.GString(tt)
		defer s.Free()
		C.gtk_entry_set_icon_tooltip_text(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition),
			(*C.gchar)(s.GetPtr()))
	}
}

func (self *Entry) GetIconTooltipText(gtk_EntryIconPosition int) string {
	s := C.gtk_entry_get_icon_tooltip_text(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition))

	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}

func (self *Entry) SetIconTooltipMarkup(gtk_EntryIconPosition int, tooltip interface{}) {
	if tooltip == nil {
		C.gtk_entry_set_icon_tooltip_markup(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition), nil)
		return
	}

	if tt, ok := tooltip.(string); ok {
		s := gobject.GString(tt)
		defer s.Free()
		C.gtk_entry_set_icon_tooltip_markup(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition),
			(*C.gchar)(s.GetPtr()))
	}
}

func (self *Entry) GetIconTooltipMarkup(gtk_EntryIconPosition int) string {
	s := C.gtk_entry_get_icon_tooltip_markup(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition))

	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}

//TODO: gtk_entry_set_icon_drag_source
//TODO: gtk_entry_get_current_icon_drag_source

func (self *Entry) GetIconArea(gtk_EntryIconPosition int) gdk3.Rectangle {
	var crec C.GdkRectangle
	C.gtk_entry_get_icon_area(self.object, C.GtkEntryIconPosition(gtk_EntryIconPosition),
		&crec)

	if rec, err := gobject.ConvertToGo(unsafe.Pointer(&crec), gdk3.GdkType.RECTANGLE); err == nil {
		return rec.(gdk3.Rectangle)
	}

	return gdk3.Rectangle{}
}

//////////////////////////////
// END GtkEntry
////////////////////////////// }}}

// GtkEntryCompletion {{{
//////////////////////////////

// GtkEntryCompletion type
type EntryCompletion struct {
	object *C.GtkEntryCompletion
}

func NewEntryCompletion() *EntryCompletion {
	ec := &EntryCompletion{}
	ec.object = C.gtk_entry_completion_new()

	if gobject.IsObjectFloating(ec) {
		gobject.RefSink(ec)
	}
	entryCompletionFinalizer(ec)

	return ec
}

//TODO: gtk_entry_completion_new_with_area

// Clear EntryCompletion struct when it goes out of reach
func entryCompletionFinalizer(ec *EntryCompletion) {
	runtime.SetFinalizer(ec, func(ec *EntryCompletion) { gobject.Unref(ec) })
}

// Conversion funcs
func newEntryCompletionFromNative(obj unsafe.Pointer) interface{} {
	ec := &EntryCompletion{}
	ec.object = C.to_GtkEntryCompletion(obj)

	if gobject.IsObjectFloating(ec) {
		gobject.RefSink(ec)
	} else {
		gobject.Ref(ec)
	}
	entryCompletionFinalizer(ec)

	return ec
}

func nativeFromEntryCompletion(ec interface{}) *gobject.GValue {
	entryCompletion, ok := ec.(*EntryCompletion)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ENTRY_COMPLETION, entryCompletion.ToNative())
		return gv
	}
	return nil
}

// To be Object-like
func (self EntryCompletion) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self EntryCompletion) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self EntryCompletion) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self EntryCompletion) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// EntryCompletion interface

func (self *EntryCompletion) GetEntry() WidgetLike {
	w := C.gtk_entry_completion_get_entry(self.object)

	if w == nil {
		return nil
	}

	widget, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return widget.(WidgetLike)
	}
	return nil
}

func (self *EntryCompletion) SetModel(model TreeModelLike) {
	var m *C.GtkTreeModel = nil

	if model != nil {
		m = model.ITreeModel().object
	}

	C.gtk_entry_completion_set_model(self.object, m)
}

func (self *EntryCompletion) GetModel() TreeModelLike {
	m := C.gtk_entry_completion_get_model(self.object)

	if m == nil {
		return nil
	}

	model, err := gobject.ConvertToGo(unsafe.Pointer(m))

	if err == nil {
		return model.(TreeModelLike)
	}
	return nil
}

func (self *EntryCompletion) SetMatchFunc(f interface{}, data ...interface{}) {
	clFunc, id := gobject.CreateCustomClosure(f, data...)

	if clFunc == nil {
		return
	}

	// Add closure to local closure map
	addToClosures(id, clFunc)

	// Register closure to C
	C._gtk_entry_completion_set_match_func(self.object, C.gint64(id))

	// Add WeakRef, so that when object is destroyed, closure func removes
	// our closure from local closure map
	gobject.WeakRef(self, func() { removeFromClosures(id) })
}

func (self *EntryCompletion) SetMinimumKeyLength(length int) {
	C.gtk_entry_completion_set_minimum_key_length(self.object, C.gint(length))
}

func (self *EntryCompletion) GetMinimumKeyLength() int {
	return int(C.gtk_entry_completion_get_minimum_key_length(self.object))
}

func (self *EntryCompletion) Complete() {
	C.gtk_entry_completion_complete(self.object)
}

func (self *EntryCompletion) GetCompletionPrefix() string {
	s := C.gtk_entry_completion_get_completion_prefix(self.object)

	if s == nil {
		return ""
	}

	return gobject.GoString(unsafe.Pointer(s))
}

func (self *EntryCompletion) InsertPrefix() {
	C.gtk_entry_completion_insert_prefix(self.object)
}

func (self *EntryCompletion) InsertActionText(index int, text string) {
	s := gobject.GString(text)
	defer s.Free()
	C.gtk_entry_completion_insert_action_text(self.object, C.gint(index), (*C.gchar)(s.GetPtr()))
}

func (self *EntryCompletion) InsertActionMarkup(index int, markup string) {
	s := gobject.GString(markup)
	defer s.Free()
	C.gtk_entry_completion_insert_action_markup(self.object, C.gint(index), (*C.gchar)(s.GetPtr()))
}

func (self *EntryCompletion) DeleteAction(index int) {
	C.gtk_entry_completion_delete_action(self.object, C.gint(index))
}

func (self *EntryCompletion) SetTextColumn(column int) {
	C.gtk_entry_completion_set_text_column(self.object, C.gint(column))
}

func (self *EntryCompletion) GetTextColumn() int {
	return int(C.gtk_entry_completion_get_text_column(self.object))
}

func (self *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	b := gobject.GBool(inlineCompletion)
	defer b.Free()
	C.gtk_entry_completion_set_inline_completion(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *EntryCompletion) GetInlineCompletion() bool {
	b := C.gtk_entry_completion_get_inline_completion(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	b := gobject.GBool(inlineSelection)
	defer b.Free()
	C.gtk_entry_completion_set_inline_selection(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *EntryCompletion) GetInlineSelection() bool {
	b := C.gtk_entry_completion_get_inline_selection(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	b := gobject.GBool(popupCompletion)
	defer b.Free()
	C.gtk_entry_completion_set_popup_completion(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *EntryCompletion) GetPopupCompletion() bool {
	b := C.gtk_entry_completion_get_popup_completion(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	b := gobject.GBool(popupSetWidth)
	defer b.Free()
	C.gtk_entry_completion_set_popup_set_width(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *EntryCompletion) GetPopupSetWidth() bool {
	b := C.gtk_entry_completion_get_popup_set_width(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	b := gobject.GBool(popupSingleMatch)
	defer b.Free()
	C.gtk_entry_completion_set_popup_single_match(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *EntryCompletion) GetPopupSingleMatch() bool {
	b := C.gtk_entry_completion_get_popup_single_match(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// END GtkEntryCompletion
////////////////////////////// }}}

// GtkScale {{{
//////////////////////////////

// GtkScale type
type Scale struct {
	object *C.GtkScale
	*Range
}

func NewScale(gtk_Orientation int, adjustment *Adjustment) *Scale {
	var adj *C.GtkAdjustment = nil
	if adjustment != nil {
		adj = adjustment.object
	}

	s := &Scale{}
	o := C.gtk_scale_new(C.GtkOrientation(gtk_Orientation), adj)
	s.object = C.to_GtkScale(unsafe.Pointer(o))

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	}
	s.Range = newRangeFromNative(unsafe.Pointer(o)).(*Range)
	scaleFinalizer(s)

	return s
}

func NewScaleWithRange(gtk_Orientation int, min, max, step float64) *Scale {
	s := &Scale{}
	o := C.gtk_scale_new_with_range(C.GtkOrientation(gtk_Orientation),
		C.gdouble(min), C.gdouble(max), C.gdouble(step))
	s.object = C.to_GtkScale(unsafe.Pointer(o))

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	}
	s.Range = newRangeFromNative(unsafe.Pointer(o)).(*Range)
	scaleFinalizer(s)

	return s
}

func NewHScale(adjustment *Adjustment) *Scale {
	return NewScale(GtkOrientation.HORIZONTAL, adjustment)
}

func NewVScale(adjustment *Adjustment) *Scale {
	return NewScale(GtkOrientation.VERTICAL, adjustment)
}

func NewHScaleWithRange(min, max, step float64) *Scale {
	return NewScaleWithRange(GtkOrientation.HORIZONTAL, min, max, step)
}

func NewVScaleWithRange(min, max, step float64) *Scale {
	return NewScaleWithRange(GtkOrientation.VERTICAL, min, max, step)
}

// Clear Scale struct when it goes out of reach
func scaleFinalizer(s *Scale) {
	runtime.SetFinalizer(s, func(s *Scale) { gobject.Unref(s) })
}

// Conversion function for gobject registration map
func newScaleFromNative(obj unsafe.Pointer) interface{} {
	s := &Scale{}
	s.object = C.to_GtkScale(obj)

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	} else {
		gobject.Ref(s)
	}
	s.Range = newRangeFromNative(obj).(*Range)
	scaleFinalizer(s)

	return s
}

func nativeFromScale(s interface{}) *gobject.GValue {
	scale, ok := s.(*Scale)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SCALE, scale.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self Scale) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Scale) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Scale) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Scale) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Range-like
func (self Scale) R() *Range {
	return self.Range
}

// Scale interface

func (self *Scale) SetDigits(digits int) {
	C.gtk_scale_set_digits(self.object, C.gint(digits))
}

func (self *Scale) SetDrawValue(drawValue bool) {
	b := gobject.GBool(drawValue)
	defer b.Free()
	C.gtk_scale_set_draw_value(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Scale) SetValuePos(gtk_Position int) {
	C.gtk_scale_set_value_pos(self.object, C.GtkPositionType(gtk_Position))
}

func (self *Scale) GetDigits() int {
	return int(C.gtk_scale_get_digits(self.object))
}

func (self *Scale) GetDrawValue() bool {
	b := C.gtk_scale_get_draw_value(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Scale) GetValuePos() int {
	return int(C.gtk_scale_get_value_pos(self.object))
}

//TODO: gtk_scale_get_layout

func (self *Scale) GetLayoutOffsets() (x, y int) {
	var cx, cy C.gint
	C.gtk_scale_get_layout_offsets(self.object, &cx, &cy)
	x = int(cx)
	y = int(cy)
	return
}

func (self *Scale) AddMark(value int64, gtk_Position int, markup interface{}) {
	if markup == nil {
		C.gtk_scale_add_mark(self.object, C.gdouble(value), C.GtkPositionType(gtk_Position), nil)
		return
	}

	if m, ok := markup.(string); ok {
		s := gobject.GString(m)
		defer s.Free()
		C.gtk_scale_add_mark(self.object, C.gdouble(value), C.GtkPositionType(gtk_Position), (*C.gchar)(s.GetPtr()))
	}
}

func (self *Scale) ClearMarks() {
	C.gtk_scale_clear_marks(self.object)
}

//////////////////////////////
// END GtkScale
////////////////////////////// }}}

// GtkSpinButton {{{
//////////////////////////////

// GtkSpinButton type
type SpinButton struct {
	object *C.GtkSpinButton
	*Entry
}

func NewSpinButton(adjustment *Adjustment, climbRate float64, digits uint) *SpinButton {
	var adj *C.GtkAdjustment = nil
	if adjustment != nil {
		adj = adjustment.object
	}

	s := &SpinButton{}
	o := C.gtk_spin_button_new(adj, C.gdouble(climbRate), C.guint(digits))
	s.object = C.to_GtkSpinButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	}
	s.Entry = newEntryFromNative(unsafe.Pointer(o)).(*Entry)
	spinButtonFinalizer(s)

	return s
}

func NewSpinButtonWithRange(min, max, step float64) *SpinButton {
	s := &SpinButton{}
	o := C.gtk_spin_button_new_with_range(C.gdouble(min), C.gdouble(max), C.gdouble(step))
	s.object = C.to_GtkSpinButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	}
	s.Entry = newEntryFromNative(unsafe.Pointer(o)).(*Entry)
	spinButtonFinalizer(s)

	return s
}

// Clear SpinButton struct when it goes out of reach
func spinButtonFinalizer(s *SpinButton) {
	runtime.SetFinalizer(s, func(s *SpinButton) { gobject.Unref(s) })
}

// Conversion function for gobject registration map
func newSpinButtonFromNative(obj unsafe.Pointer) interface{} {
	s := &SpinButton{}
	s.object = C.to_GtkSpinButton(obj)

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	} else {
		gobject.Ref(s)
	}
	s.Entry = newEntryFromNative(obj).(*Entry)
	spinButtonFinalizer(s)

	return s
}

func nativeFromSpinButton(s interface{}) *gobject.GValue {
	spin, ok := s.(*SpinButton)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SPIN_BUTTON, spin.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self SpinButton) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self SpinButton) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self SpinButton) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self SpinButton) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Entry-like
func (self SpinButton) E() *Entry {
	return self.Entry
}

// SpinButton interface

func (self *SpinButton) Configure(adj *Adjustment, climbRate float64, digits uint) {
	var cadj *C.GtkAdjustment = nil
	if adj != nil {
		cadj = adj.object
	}

	C.gtk_spin_button_configure(self.object, cadj, C.gdouble(climbRate), C.guint(digits))
}

func (self *SpinButton) SetAdjustment(adj *Adjustment) {
	C.gtk_spin_button_set_adjustment(self.object, adj.object)
}

func (self *SpinButton) GetAdjustment() *Adjustment {
	a := C.gtk_spin_button_get_adjustment(self.object)

	if a != nil {
		if adj, err := gobject.ConvertToGo(unsafe.Pointer(a)); err == nil {
			return adj.(*Adjustment)
		}
	}
	return nil
}

func (self *SpinButton) SetDigits(digits uint) {
	C.gtk_spin_button_set_digits(self.object, C.guint(digits))
}

func (self *SpinButton) SetIncrements(step, page float64) {
	C.gtk_spin_button_set_increments(self.object, C.gdouble(step), C.gdouble(page))
}

func (self *SpinButton) SetRange(min, max float64) {
	C.gtk_spin_button_set_range(self.object, C.gdouble(min), C.gdouble(max))
}

func (self *SpinButton) GetValueAsInt() int {
	return int(C.gtk_spin_button_get_value_as_int(self.object))
}

func (self *SpinButton) SetValue(value float64) {
	C.gtk_spin_button_set_value(self.object, C.gdouble(value))
}

func (self *SpinButton) SetUpdatePolicy(gtk_SpinButtonUpdatePolicy int) {
	C.gtk_spin_button_set_update_policy(self.object,
		C.GtkSpinButtonUpdatePolicy(gtk_SpinButtonUpdatePolicy))
}

func (self *SpinButton) SetNumeric(numeric bool) {
	b := gobject.GBool(numeric)
	defer b.Free()
	C.gtk_spin_button_set_numeric(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *SpinButton) Spin(gtk_SpinType int, increment float64) {
	C.gtk_spin_button_spin(self.object, C.GtkSpinType(gtk_SpinType), C.gdouble(increment))
}

func (self *SpinButton) SetWrap(wrap bool) {
	b := gobject.GBool(wrap)
	defer b.Free()
	C.gtk_spin_button_set_wrap(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *SpinButton) SetSnapToTicks(snapToTicks bool) {
	b := gobject.GBool(snapToTicks)
	defer b.Free()
	C.gtk_spin_button_set_snap_to_ticks(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *SpinButton) Update() {
	C.gtk_spin_button_update(self.object)
}

func (self *SpinButton) GetDigits() uint {
	return uint(C.gtk_spin_button_get_digits(self.object))
}

func (self *SpinButton) GetIncrements() (step, page float64) {
	var f, p C.gdouble
	C.gtk_spin_button_get_increments(self.object, &f, &p)
	step = float64(f)
	page = float64(p)
	return
}

func (self *SpinButton) GetNumeric() bool {
	b := C.gtk_spin_button_get_numeric(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *SpinButton) GetRange() (min, max float64) {
	var cmin, cmax C.gdouble
	C.gtk_spin_button_get_range(self.object, &cmin, &cmax)
	min = float64(cmin)
	max = float64(cmax)
	return
}

func (self *SpinButton) GetSnapToTicks() bool {
	b := C.gtk_spin_button_get_snap_to_ticks(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *SpinButton) GetUpdatePolicy() int {
	return int(C.gtk_spin_button_get_update_policy(self.object))
}

func (self *SpinButton) GetValue() float64 {
	return float64(C.gtk_spin_button_get_value(self.object))
}

func (self *SpinButton) GetWrap() bool {
	b := C.gtk_spin_button_get_wrap(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// END GtkSpinButton
////////////////////////////// }}}

// End Numeric/Text Data Entry }}}

// Multiline Text Editor {{{

// GtkTextTag {{{
//////////////////////////////

// GtkTextTag type
type TextTag struct {
	object *C.GtkTextTag
}

func NewTextTag(name string) *TextTag {
	t := &TextTag{}
	s := gobject.GString(name)
	defer s.Free()
	o := C.gtk_text_tag_new((*C.gchar)(s.GetPtr()))
	t.object = C.to_GtkTextTag(unsafe.Pointer(o))

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	}
	textTagFinalizer(t)

	return t
}

// Clear TextTag struct when it goes out of reach
func textTagFinalizer(t *TextTag) {
	runtime.SetFinalizer(t, func(t *TextTag) { gobject.Unref(t) })
}

// Conversion funcs
func newTextTagFromNative(obj unsafe.Pointer) interface{} {
	t := &TextTag{}
	t.object = C.to_GtkTextTag(obj)

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	} else {
		gobject.Ref(t)
	}
	textTagFinalizer(t)

	return t
}

func nativeFromTextTag(t interface{}) *gobject.GValue {
	ttt, ok := t.(*TextTag)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TEXT_TAG, ttt.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TextTag) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TextTag) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TextTag) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TextTag) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// TextTag interface

func (self *TextTag) GetPriority() int {
	return int(C.gtk_text_tag_get_priority(self.object))
}

func (self *TextTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority(self.object, C.gint(priority))
}

//TODO: gtk_text_tag_event

//////////////////////////////
// End GtkTextTag
////////////////////////////// }}}

// GtkTextTagTable {{{
//////////////////////////////

// GtkTextTagTable type
type TextTagTable struct {
	object *C.GtkTextTagTable
}

func NewTextTagTable() *TextTagTable {
	t := &TextTagTable{}
	o := C.gtk_text_tag_table_new()
	t.object = C.to_GtkTextTagTable(unsafe.Pointer(o))

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	}
	textTagTableFinalizer(t)

	return t
}

// Clear TextTagTable struct when it goes out of reach
func textTagTableFinalizer(t *TextTagTable) {
	runtime.SetFinalizer(t, func(t *TextTagTable) { gobject.Unref(t) })
}

// Conversion funcs
func newTextTagTableFromNative(obj unsafe.Pointer) interface{} {
	t := &TextTagTable{}
	t.object = C.to_GtkTextTagTable(obj)

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	} else {
		gobject.Ref(t)
	}
	textTagTableFinalizer(t)

	return t
}

func nativeFromTextTagTable(t interface{}) *gobject.GValue {
	ttt, ok := t.(*TextTagTable)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TEXT_TAG_TABLE, ttt.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TextTagTable) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TextTagTable) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TextTagTable) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TextTagTable) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// TextTagTable interface

func (self *TextTagTable) Add(tag *TextTag) {
	C.gtk_text_tag_table_add(self.object, tag.object)
}

func (self *TextTagTable) Remove(tag *TextTag) {
	C.gtk_text_tag_table_remove(self.object, tag.object)
}

func (self *TextTagTable) Lookup(name string) *TextTag {
	s := gobject.GString(name)
	defer s.Free()
	t := C.gtk_text_tag_table_lookup(self.object, (*C.gchar)(s.GetPtr()))

	if t != nil {
		tag, err := gobject.ConvertToGo(unsafe.Pointer(t))
		if err == nil {
			return tag.(*TextTag)
		}
	}
	return nil
}

func (self *TextTagTable) Foreach(f interface{}, data ...interface{}) {
	c, id := gobject.CreateCustomClosure(f, data...)
	// Add f to closures
	_closures[id] = c
	// Call C foreach
	C._gtk_text_tag_table_foreach(self.object, C.gint64(id))

	// Remove closure from _closures
	delete(_closures, id)
}

func (self *TextTagTable) GetSize() int {
	return int(C.gtk_text_tag_table_get_size(self.object))
}

//////////////////////////////
// End GtkTextTagTable
////////////////////////////// }}}

// GtkTextMark {{{
//////////////////////////////

// GtkTextMark type
type TextMark struct {
	object *C.GtkTextMark
}

func NewTextMark(name interface{}, leftGravity bool) *TextMark {
	var markName *C.gchar = nil
	if name != nil {
		if n, ok := name.(string); ok {
			s := gobject.GString(n)
			defer s.Free()
			markName = (*C.gchar)(s.GetPtr())
		}
	}
	b := gobject.GBool(leftGravity)
	defer b.Free()

	t := &TextMark{}
	o := C.gtk_text_mark_new(markName, *((*C.gboolean)(b.GetPtr())))
	t.object = C.to_GtkTextMark(unsafe.Pointer(o))

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	}
	textMarkFinalizer(t)

	return t
}

// Clear TextMark struct when it goes out of reach
func textMarkFinalizer(t *TextMark) {
	runtime.SetFinalizer(t, func(t *TextMark) { gobject.Unref(t) })
}

// Conversion funcs
func newTextMarkFromNative(obj unsafe.Pointer) interface{} {
	t := &TextMark{}
	t.object = C.to_GtkTextMark(obj)

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	} else {
		gobject.Ref(t)
	}
	textMarkFinalizer(t)

	return t
}

func nativeFromTextMark(t interface{}) *gobject.GValue {
	tm, ok := t.(*TextMark)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TEXT_MARK, tm.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TextMark) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TextMark) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TextMark) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TextMark) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// TextMark interface

func (self *TextMark) SetVisible(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_text_mark_set_visible(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TextMark) GetVisible() bool {
	b := C.gtk_text_mark_get_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextMark) GetDeleted() bool {
	b := C.gtk_text_mark_get_deleted(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextMark) GetName() string {
	n := C.gtk_text_mark_get_name(self.object)
	if n != nil {
		return gobject.GoString(unsafe.Pointer(n))
	}
	return ""
}

func (self *TextMark) GetBuffer() *TextBuffer {
	b := C.gtk_text_mark_get_buffer(self.object)

	if b == nil {
		return nil
	}

	if buf, err := gobject.ConvertToGo(unsafe.Pointer(b)); err == nil {
		return buf.(*TextBuffer)
	}
	return nil
}

func (self *TextMark) GetLeftGravity() bool {
	b := C.gtk_text_mark_get_left_gravity(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// End GtkTextMark
////////////////////////////// }}}

// GtkTextIter {{{
//////////////////////////////

// GtkTextIter type
type TextIter struct {
	object C.GtkTextIter
}

// Conversion funcs
func newTextIterFromNative(obj unsafe.Pointer) interface{} {
	ti := &TextIter{*((*C.GtkTextIter)(obj))}
	return ti
}

func nativeFromTextIter(ti interface{}) *gobject.GValue {
	t, ok := ti.(*TextIter)
	if ok {
		it := C.gtk_text_iter_copy(&t.object)
		// New GValue object becomes the owner of it so we don't have to free it
		gv := gobject.CreateCGValue(GtkType.TEXT_ITER, unsafe.Pointer(it))
		return gv
	}
	return nil
}

// TextIter is boxed-like
func (self TextIter) GetBoxedType() gobject.GType {
	return GtkType.TEXT_ITER
}

// TextIter interface

func (self *TextIter) GetBuffer() *TextBuffer {
	b := C.gtk_text_iter_get_buffer(&self.object)

	if b == nil {
		return nil
	}

	if buf, err := gobject.ConvertToGo(unsafe.Pointer(b)); err == nil {
		return buf.(*TextBuffer)
	}
	return nil
}

func (self *TextIter) Assign(iter *TextIter) {
	iter.object = self.object
}

func (self *TextIter) GetOffset() int {
	return int(C.gtk_text_iter_get_offset(&self.object))
}

func (self *TextIter) GetLine() int {
	return int(C.gtk_text_iter_get_line(&self.object))
}

func (self *TextIter) GetLineOffset() int {
	return int(C.gtk_text_iter_get_line_offset(&self.object))
}

func (self *TextIter) GetLineIndex() int {
	return int(C.gtk_text_iter_get_line_index(&self.object))
}

func (self *TextIter) GetVisibleLineIndex() int {
	return int(C.gtk_text_iter_get_visible_line_index(&self.object))
}

func (self *TextIter) GetVisibleLineOffset() int {
	return int(C.gtk_text_iter_get_visible_line_offset(&self.object))
}

func (self *TextIter) GetChar() rune {
	return rune(C.gtk_text_iter_get_char(&self.object))
}

func (self *TextIter) GetSlice(end *TextIter) string {
	s := C.gtk_text_iter_get_slice(&self.object, &end.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *TextIter) GetText(end *TextIter) string {
	s := C.gtk_text_iter_get_text(&self.object, &end.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *TextIter) GetVisibleSlice(end *TextIter) string {
	s := C.gtk_text_iter_get_visible_slice(&self.object, &end.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *TextIter) GetVisibleText(end *TextIter) string {
	s := C.gtk_text_iter_get_visible_text(&self.object, &end.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *TextIter) GetPixbuf() *gdkpixbuf.Pixbuf {
	p := C.gtk_text_iter_get_pixbuf(&self.object)
	if p != nil {
		pix, err := gobject.ConvertToGo(unsafe.Pointer(p))
		if err == nil {
			return pix.(*gdkpixbuf.Pixbuf)
		}
	}
	return nil
}

func (self *TextIter) GetMarks() *glib.GSList {
	sl := C.gtk_text_iter_get_marks(&self.object)
	goList := glib.NewGSListFromNative(unsafe.Pointer(sl))

	// Free the list but not elements
	goList.GC_Free = true
	goList.GC_FreeFull = false
	glib.GSListFinalizer(goList)

	// List contains TextMarks so we setup converter func for them
	goList.ConversionFunc = newTextMarkFromNative

	return goList
}

func (self *TextIter) GetToggledTags(toggledOn bool) *glib.GSList {
	b := gobject.GBool(toggledOn)
	defer b.Free()
	tags := C.gtk_text_iter_get_toggled_tags(&self.object, *((*C.gboolean)(b.GetPtr())))
	goList := glib.NewGSListFromNative(unsafe.Pointer(tags))

	// Free the list but not elements
	goList.GC_Free = true
	goList.GC_FreeFull = false
	glib.GSListFinalizer(goList)

	// List contains TextTags so we setup converter func for them
	goList.ConversionFunc = newTextTagFromNative

	return goList
}

func (self *TextIter) GetChildAnchor() *TextChildAnchor {
	o := C.gtk_text_iter_get_child_anchor(&self.object)

	if o == nil {
		return nil
	}

	if tca, err := gobject.ConvertToGo(unsafe.Pointer(o)); err == nil {
		return tca.(*TextChildAnchor)
	}
	return nil
}

func (self *TextIter) BeginsTag(tag *TextTag) bool {
	var b C.gboolean
	if tag != nil {
		b = C.gtk_text_iter_begins_tag(&self.object, tag.object)
	} else {
		b = C.gtk_text_iter_begins_tag(&self.object, nil)
	}
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) EndsTag(tag *TextTag) bool {
	var b C.gboolean
	if tag != nil {
		b = C.gtk_text_iter_ends_tag(&self.object, tag.object)
	} else {
		b = C.gtk_text_iter_ends_tag(&self.object, nil)
	}
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) TogglesTag(tag *TextTag) bool {
	var b C.gboolean
	if tag != nil {
		b = C.gtk_text_iter_toggles_tag(&self.object, tag.object)
	} else {
		b = C.gtk_text_iter_toggles_tag(&self.object, nil)
	}
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) HasTag(tag *TextTag) bool {
	if tag == nil {
		return false
	}
	b := C.gtk_text_iter_has_tag(&self.object, tag.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) GetTags() *glib.GSList {
	tags := C.gtk_text_iter_get_tags(&self.object)
	goList := glib.NewGSListFromNative(unsafe.Pointer(tags))

	// Free the list but not elements
	goList.GC_Free = true
	goList.GC_FreeFull = false
	glib.GSListFinalizer(goList)

	// List contains TextTags so we setup converter func for them
	goList.ConversionFunc = newTextTagFromNative

	return goList
}

func (self *TextIter) Editable(defaultSetting bool) bool {
	b := gobject.GBool(defaultSetting)
	defer b.Free()

	res := C.gtk_text_iter_editable(&self.object, *((*C.gboolean)(b.GetPtr())))
	return gobject.GoBool(unsafe.Pointer(&res))
}

func (self *TextIter) CanInsert(defaultEditability bool) bool {
	b := gobject.GBool(defaultEditability)
	defer b.Free()

	res := C.gtk_text_iter_can_insert(&self.object, *((*C.gboolean)(b.GetPtr())))
	return gobject.GoBool(unsafe.Pointer(&res))
}

func (self *TextIter) StartsWord() bool {
	b := C.gtk_text_iter_starts_word(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) EndsWord() bool {
	b := C.gtk_text_iter_ends_word(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) InsideWord() bool {
	b := C.gtk_text_iter_inside_word(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) StartsLine() bool {
	b := C.gtk_text_iter_starts_line(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) EndsLine() bool {
	b := C.gtk_text_iter_ends_line(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) StartsSentence() bool {
	b := C.gtk_text_iter_starts_sentence(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) EndsSentence() bool {
	b := C.gtk_text_iter_ends_sentence(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) InsideSentence() bool {
	b := C.gtk_text_iter_inside_sentence(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) IsCursorPosition() bool {
	b := C.gtk_text_iter_is_cursor_position(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) GetCharsInLine() int {
	return int(C.gtk_text_iter_get_chars_in_line(&self.object))
}

func (self *TextIter) GetBytesInLine() int {
	return int(C.gtk_text_iter_get_bytes_in_line(&self.object))
}

//TODO: gtk_text_iter_get_attributes
//TODO: gtk_text_iter_get_language

func (self *TextIter) IsEnd() bool {
	b := C.gtk_text_iter_is_end(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) IsStart() bool {
	b := C.gtk_text_iter_is_start(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardChar() bool {
	b := C.gtk_text_iter_forward_char(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardChar() bool {
	b := C.gtk_text_iter_forward_char(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardChars(count int) bool {
	b := C.gtk_text_iter_forward_chars(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardChars(count int) bool {
	b := C.gtk_text_iter_backward_chars(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardLine() bool {
	b := C.gtk_text_iter_forward_line(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardLine() bool {
	b := C.gtk_text_iter_backward_line(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardLines(count int) bool {
	b := C.gtk_text_iter_forward_lines(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardLines(count int) bool {
	b := C.gtk_text_iter_backward_lines(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardWordEnds(count int) bool {
	b := C.gtk_text_iter_forward_word_ends(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardWordStarts(count int) bool {
	b := C.gtk_text_iter_backward_word_starts(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardWordEnd() bool {
	b := C.gtk_text_iter_forward_word_end(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardWordStart() bool {
	b := C.gtk_text_iter_backward_word_start(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardCursorPosition() bool {
	b := C.gtk_text_iter_forward_cursor_position(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardCursorPosition() bool {
	b := C.gtk_text_iter_backward_cursor_position(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardCursorPositions(count int) bool {
	b := C.gtk_text_iter_forward_cursor_positions(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardCursorPositions(count int) bool {
	b := C.gtk_text_iter_backward_cursor_positions(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardSentenceStart() bool {
	b := C.gtk_text_iter_backward_sentence_start(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardSentenceStarts(count int) bool {
	b := C.gtk_text_iter_backward_sentence_starts(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardSentenceEnd() bool {
	b := C.gtk_text_iter_forward_sentence_end(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardSentenceEnds(count int) bool {
	b := C.gtk_text_iter_forward_sentence_ends(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardVisibleWordEnds(count int) bool {
	b := C.gtk_text_iter_forward_visible_word_ends(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardVisibleWordStarts(count int) bool {
	b := C.gtk_text_iter_backward_visible_word_starts(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardVisibleWordEnd() bool {
	b := C.gtk_text_iter_forward_visible_word_end(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardVisibleWordStart() bool {
	b := C.gtk_text_iter_backward_visible_word_start(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardVisibleCursorPosition() bool {
	b := C.gtk_text_iter_forward_visible_cursor_position(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardVisibleCursorPosition() bool {
	b := C.gtk_text_iter_backward_visible_cursor_position(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardVisibleCursorPositions(count int) bool {
	b := C.gtk_text_iter_forward_visible_cursor_positions(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardVisibleCursorPositions(count int) bool {
	b := C.gtk_text_iter_backward_visible_cursor_positions(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardVisibleLine() bool {
	b := C.gtk_text_iter_forward_visible_line(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardVisibleLine() bool {
	b := C.gtk_text_iter_backward_visible_line(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardVisibleLines(count int) bool {
	b := C.gtk_text_iter_forward_visible_lines(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardVisibleLines(count int) bool {
	b := C.gtk_text_iter_backward_visible_lines(&self.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) SetOffset(charOffset int) {
	C.gtk_text_iter_set_offset(&self.object, C.gint(charOffset))
}

func (self *TextIter) SetLine(lineNumber int) {
	C.gtk_text_iter_set_line(&self.object, C.gint(lineNumber))
}

func (self *TextIter) SetLineOffset(charOnLine int) {
	C.gtk_text_iter_set_line_offset(&self.object, C.gint(charOnLine))
}

func (self *TextIter) SetLineIndex(bytesOnLine int) {
	C.gtk_text_iter_set_line_index(&self.object, C.gint(bytesOnLine))
}

func (self *TextIter) SetVisibleLineIndex(bytesOnLine int) {
	C.gtk_text_iter_set_visible_line_index(&self.object, C.gint(bytesOnLine))
}

func (self *TextIter) SetVisibleLineOffset(charOnLine int) {
	C.gtk_text_iter_set_visible_line_offset(&self.object, C.gint(charOnLine))
}

func (self *TextIter) ForwardToEnd() {
	C.gtk_text_iter_forward_to_end(&self.object)
}

func (self *TextIter) ForwardToLineEnd() bool {
	b := C.gtk_text_iter_forward_to_line_end(&self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardToTagToggle(tag *TextTag) bool {
	b := C.gtk_text_iter_forward_to_tag_toggle(&self.object, tag.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardToTagToggle(tag *TextTag) bool {
	b := C.gtk_text_iter_backward_to_tag_toggle(&self.object, tag.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardFindChar(limit *TextIter, f interface{}, data ...interface{}) bool {
	call, id := gobject.CreateCustomClosure(f, data...)
	_closures[id] = call
	b := C._gtk_text_iter_forward_find_char(&self.object, &limit.object, C.gint64(id))

	delete(_closures, id)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardFindChar(limit *TextIter, f interface{}, data ...interface{}) bool {
	call, id := gobject.CreateCustomClosure(f, data...)
	_closures[id] = call
	b := C._gtk_text_iter_backward_find_char(&self.object, &limit.object, C.gint64(id))

	delete(_closures, id)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) ForwardSearch(str string, gtk_TextSearchFlags int, matchStart, matchEnd, limit *TextIter) bool {
	var s, e, l *C.GtkTextIter = nil, nil, nil

	if matchStart != nil {
		s = &matchStart.object
	}

	if matchEnd != nil {
		e = &matchEnd.object
	}

	if limit != nil {
		l = &limit.object
	}

	cs := gobject.GString(str)
	defer cs.Free()

	b := C.gtk_text_iter_forward_search(&self.object, (*C.gchar)(cs.GetPtr()), C.GtkTextSearchFlags(gtk_TextSearchFlags), s, e, l)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) BackwardSearch(str string, gtk_TextSearchFlags int, matchStart, matchEnd, limit *TextIter) bool {
	var s, e, l *C.GtkTextIter = nil, nil, nil

	if matchStart != nil {
		s = &matchStart.object
	}

	if matchEnd != nil {
		e = &matchEnd.object
	}

	if limit != nil {
		l = &limit.object
	}

	cs := gobject.GString(str)
	defer cs.Free()

	b := C.gtk_text_iter_backward_search(&self.object, (*C.gchar)(cs.GetPtr()), C.GtkTextSearchFlags(gtk_TextSearchFlags), s, e, l)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) Equal(rhs *TextIter) bool {
	b := C.gtk_text_iter_equal(&self.object, &rhs.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) Compare(rhs *TextIter) bool {
	b := C.gtk_text_iter_compare(&self.object, &rhs.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) InRange(start, end *TextIter) bool {
	b := C.gtk_text_iter_in_range(&self.object, &start.object, &end.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextIter) Order(second *TextIter) {
	C.gtk_text_iter_order(&self.object, &second.object)
}

//////////////////////////////
// End GtkTextIter
////////////////////////////// }}}

// GtkTextBuffer {{{
//////////////////////////////

// GtkTextBuffer type
type TextBuffer struct {
	object *C.GtkTextBuffer
}

func NewTextBuffer(table *TextTagTable) *TextBuffer {
	var ttt *C.GtkTextTagTable = nil
	if table != nil {
		ttt = table.object
	}
	t := &TextBuffer{}
	t.object = C.gtk_text_buffer_new(ttt)

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	}
	textBufferFinalizer(t)

	return t
}

// Clear TextBuffer struct when it goes out of reach
func textBufferFinalizer(t *TextBuffer) {
	runtime.SetFinalizer(t, func(t *TextBuffer) { gobject.Unref(t) })
}

// Conversion funcs
func newTextBufferFromNative(obj unsafe.Pointer) interface{} {
	t := &TextBuffer{}
	t.object = C.to_GtkTextBuffer(obj)

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	} else {
		gobject.Ref(t)
	}
	textBufferFinalizer(t)

	return t
}

func nativeFromTextBuffer(t interface{}) *gobject.GValue {
	tb, ok := t.(*TextBuffer)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TEXT_TAG, tb.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TextBuffer) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TextBuffer) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TextBuffer) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TextBuffer) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// TextBuffer interface

func (self *TextBuffer) GetLineCount() int {
	return int(C.gtk_text_buffer_get_line_count(self.object))
}

func (self *TextBuffer) GetCharCount() int {
	return int(C.gtk_text_buffer_get_char_count(self.object))
}

func (self *TextBuffer) GetTagTable() *TextTagTable {
	t := C.gtk_text_buffer_get_tag_table(self.object)
	if t != nil {
		if table, err := gobject.ConvertToGo(unsafe.Pointer(t)); err == nil {
			return table.(*TextTagTable)
		}
	}
	return nil
}

func (self *TextBuffer) Insert(iter *TextIter, text string) {
	s := gobject.GString(text)
	defer s.Free()

	C.gtk_text_buffer_insert(self.object, &iter.object, (*C.gchar)(s.GetPtr()), C.gint(len(text)))
}

func (self *TextBuffer) InsertAtCursor(text string) {
	s := gobject.GString(text)
	defer s.Free()

	C.gtk_text_buffer_insert_at_cursor(self.object, (*C.gchar)(s.GetPtr()), C.gint(len(text)))
}

func (self *TextBuffer) InsertInteractive(iter *TextIter, text string, defaultEditable bool) bool {
	s := gobject.GString(text)
	defer s.Free()
	b := gobject.GBool(defaultEditable)
	defer b.Free()
	res := C.gtk_text_buffer_insert_interactive(self.object, &iter.object, (*C.gchar)(s.GetPtr()),
		C.gint(len(text)), *((*C.gboolean)(b.GetPtr())))

	return gobject.GoBool(unsafe.Pointer(&res))
}

func (self *TextBuffer) InsertInteractiveAtCursor(text string, defaultEditable bool) bool {
	s := gobject.GString(text)
	defer s.Free()
	b := gobject.GBool(defaultEditable)
	defer b.Free()
	res := C.gtk_text_buffer_insert_interactive_at_cursor(self.object, (*C.gchar)(s.GetPtr()),
		C.gint(len(text)), *((*C.gboolean)(b.GetPtr())))

	return gobject.GoBool(unsafe.Pointer(&res))
}

func (self *TextBuffer) InsertRange(iter, start, end *TextIter) {
	C.gtk_text_buffer_insert_range(self.object, &iter.object, &start.object, &end.object)
}

func (self *TextBuffer) InsertRangeInteractive(iter, start, end *TextIter, defaultEditable bool) bool {
	b := gobject.GBool(defaultEditable)
	defer b.Free()
	res := C.gtk_text_buffer_insert_range_interactive(self.object, &iter.object, &start.object,
		&end.object, *((*C.gboolean)(b.GetPtr())))

	return gobject.GoBool(unsafe.Pointer(&res))
}

func (self *TextBuffer) InsertWithTags(iter *TextIter, text string, tags ...*TextTag) {
	var startIter TextIter

	startOffset := iter.GetOffset()
	self.Insert(iter, text)
	self.GetIterAtOffset(&startIter, startOffset)

	for _, tag := range tags {
		self.ApplyTag(tag, &startIter, iter)
	}
}

func (self *TextBuffer) InsertWithTagsByName(iter *TextIter, text string, tagNames ...string) {
	var startIter TextIter

	startOffset := iter.GetOffset()
	self.Insert(iter, text)
	self.GetIterAtOffset(&startIter, startOffset)

	tagTable := self.GetTagTable()

	for _, name := range tagNames {
		tag := tagTable.Lookup(name)

		if tag == nil {
			return
		}
		self.ApplyTag(tag, &startIter, iter)
	}
}

func (self *TextBuffer) Delete(start, end *TextIter) {
	C.gtk_text_buffer_delete(self.object, &start.object, &end.object)
}

func (self *TextBuffer) DeleteInteractive(start, end *TextIter, defaultEditable bool) bool {
	b := gobject.GBool(defaultEditable)
	defer b.Free()
	res := C.gtk_text_buffer_delete_interactive(self.object, &start.object, &end.object, *((*C.gboolean)(b.GetPtr())))

	return gobject.GoBool(unsafe.Pointer(&res))
}

func (self *TextBuffer) Backspace(iter *TextIter, interactive, defaultEditable bool) bool {
	b := gobject.GBool(defaultEditable)
	defer b.Free()
	i := gobject.GBool(interactive)
	defer i.Free()
	res := C.gtk_text_buffer_backspace(self.object, &iter.object, *((*C.gboolean)(i.GetPtr())), *((*C.gboolean)(b.GetPtr())))

	return gobject.GoBool(unsafe.Pointer(&res))
}

func (self *TextBuffer) SetText(text string) {
	s := gobject.GString(text)
	defer s.Free()
	C.gtk_text_buffer_set_text(self.object, (*C.gchar)(s.GetPtr()), C.gint(len(text)))
}

func (self *TextBuffer) GetText(start, end *TextIter, includeHidden bool) string {
	b := gobject.GBool(includeHidden)
	defer b.Free()

	s := C.gtk_text_buffer_get_slice(self.object, &start.object, &end.object, *((*C.gboolean)(b.GetPtr())))
	if s == nil {
		return ""
	}

	return gobject.GoString(unsafe.Pointer(s))
}

func (self *TextBuffer) InsertPixbuf(iter *TextIter, pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_text_buffer_insert_pixbuf(self.object, &iter.object, (*C.GdkPixbuf)(pixbuf.ToNative()))
}

func (self *TextBuffer) InsertChildAnchor(iter *TextIter, anchor *TextChildAnchor) {
	if anchor == nil {
		return
	}
	C.gtk_text_buffer_insert_child_anchor(self.object, &iter.object, anchor.object)
}

func (self *TextBuffer) CreateChildAnchor(iter *TextIter) *TextChildAnchor {
	o := C.gtk_text_buffer_create_child_anchor(self.object, &iter.object)

	if o == nil {
		return nil
	}
	//	if tca, err := gobject.ConvertToGo(unsafe.Pointer(o), GtkType.TEXT_CHILD_ANCHOR); err == nil {
	if tca, err := gobject.ConvertToGo(unsafe.Pointer(o)); err == nil {
		return tca.(*TextChildAnchor)
	}
	return nil
}

func (self *TextBuffer) CreateMark(name string, where *TextIter, leftGravity bool) *TextMark {
	var s *C.gchar

	if name == "" {
		s = nil
	} else {
		gs := gobject.GString(name)
		defer gs.Free()
		s = (*C.gchar)(gs.GetPtr())
	}

	b := gobject.GBool(leftGravity)
	defer b.Free()

	t := C.gtk_text_buffer_create_mark(self.object, s, &where.object, *((*C.gboolean)(b.GetPtr())))

	if t != nil {
		if mark, err := gobject.ConvertToGo(unsafe.Pointer(t)); err == nil {
			return mark.(*TextMark)
		}
	}
	return nil
}

func (self *TextBuffer) MoveMark(mark *TextMark, where *TextIter) {
	C.gtk_text_buffer_move_mark(self.object, mark.object, &where.object)
}

func (self *TextBuffer) MoveMarkByName(name string, where *TextIter) {
	var s *C.gchar
	if name == "" {
		s = nil
	} else {
		n := gobject.GString(name)
		defer n.Free()
		s = (*C.gchar)(n.GetPtr())
	}

	C.gtk_text_buffer_move_mark_by_name(self.object, s, &where.object)
}

func (self *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	C.gtk_text_buffer_add_mark(self.object, mark.object, &where.object)
}

func (self *TextBuffer) DeleteMark(mark *TextMark) {
	C.gtk_text_buffer_delete_mark(self.object, mark.object)
}

func (self *TextBuffer) DeleteMarkByName(name string) {
	s := gobject.GString(name)
	defer s.Free()

	C.gtk_text_buffer_delete_mark_by_name(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *TextBuffer) GetMark(name string) *TextMark {
	s := gobject.GString(name)
	defer s.Free()

	m := C.gtk_text_buffer_get_mark(self.object, (*C.gchar)(s.GetPtr()))

	if m == nil {
		return nil
	}

	if mark, err := gobject.ConvertToGo(unsafe.Pointer(m)); err == nil {
		return mark.(*TextMark)
	}
	return nil
}

func (self *TextBuffer) GetInsert() *TextMark {
	m := C.gtk_text_buffer_get_insert(self.object)

	if m == nil {
		return nil
	}
	if mark, err := gobject.ConvertToGo(unsafe.Pointer(m)); err == nil {
		return mark.(*TextMark)
	}
	return nil
}

func (self *TextBuffer) GetSelectionBound() *TextMark {
	m := C.gtk_text_buffer_get_selection_bound(self.object)

	if m == nil {
		return nil
	}
	if mark, err := gobject.ConvertToGo(unsafe.Pointer(m)); err == nil {
		return mark.(*TextMark)
	}
	return nil
}

func (self *TextBuffer) GetHasSelection() bool {
	b := C.gtk_text_buffer_get_has_selection(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextBuffer) PlaceCursor(where *TextIter) {
	C.gtk_text_buffer_place_cursor(self.object, &where.object)
}

func (self *TextBuffer) SelectRange(ins, bound *TextIter) {
	C.gtk_text_buffer_select_range(self.object, &ins.object, &bound.object)
}

func (self *TextBuffer) ApplyTag(tag *TextTag, start, end *TextIter) {
	C.gtk_text_buffer_apply_tag(self.object, tag.object, &start.object, &end.object)
}

func (self *TextBuffer) RemoveTag(tag *TextTag, start, end *TextIter) {
	C.gtk_text_buffer_remove_tag(self.object, tag.object, &start.object, &end.object)
}

func (self *TextBuffer) ApplyTagByName(name string, start, end *TextIter) {
	s := gobject.GString(name)
	defer s.Free()
	C.gtk_text_buffer_apply_tag_by_name(self.object, (*C.gchar)(s.GetPtr()), &start.object, &end.object)
}

func (self *TextBuffer) RemoveTagByName(name string, start, end *TextIter) {
	s := gobject.GString(name)
	defer s.Free()
	C.gtk_text_buffer_remove_tag_by_name(self.object, (*C.gchar)(s.GetPtr()), &start.object, &end.object)
}

func (self *TextBuffer) RemoveAllTags(start, end *TextIter) {
	C.gtk_text_buffer_remove_all_tags(self.object, &start.object, &end.object)
}

func (self *TextBuffer) CreateTag(name string, properties P) *TextTag {
	if len(properties) == 0 {
		return nil
	}

	tag := NewTextTag(name)
	self.GetTagTable().Add(tag)

	tag.Set(properties)

	return tag
}

func (self *TextBuffer) GetIterAtLineOffset(iter *TextIter, lineNumber int, charOffset int) {
	C.gtk_text_buffer_get_iter_at_line_offset(self.object, &iter.object, C.gint(lineNumber), C.gint(charOffset))
}

func (self *TextBuffer) GetIterAtOffset(iter *TextIter, charOffset int) {
	C.gtk_text_buffer_get_iter_at_offset(self.object, &iter.object, C.gint(charOffset))
}

func (self *TextBuffer) GetIterAtLine(iter *TextIter, lineNumber int) {
	C.gtk_text_buffer_get_iter_at_line(self.object, &iter.object, C.gint(lineNumber))
}

func (self *TextBuffer) GetIterAtLineIndex(iter *TextIter, lineNumber int, byteIndex int) {
	C.gtk_text_buffer_get_iter_at_line_index(self.object, &iter.object, C.gint(lineNumber), C.gint(byteIndex))
}

func (self *TextBuffer) GetIterAtMark(iter *TextIter, mark *TextMark) {
	C.gtk_text_buffer_get_iter_at_mark(self.object, &iter.object, mark.object)
}

//TODO: gtk_text_buffer_get_iter_at_child_anchor

func (self *TextBuffer) GetStartIter(iter *TextIter) {
	C.gtk_text_buffer_get_start_iter(self.object, &iter.object)
}

func (self *TextBuffer) GetEndIter(iter *TextIter) {
	C.gtk_text_buffer_get_end_iter(self.object, &iter.object)
}

func (self *TextBuffer) GetBounds(start, end *TextIter) {
	C.gtk_text_buffer_get_bounds(self.object, &start.object, &end.object)
}

func (self *TextBuffer) GetModified() bool {
	b := C.gtk_text_buffer_get_modified(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextBuffer) SetModified(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_text_buffer_set_modified(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TextBuffer) DeleteSelection(interactive, defaultEditable bool) bool {
	arg1 := gobject.GBool(interactive)
	defer arg1.Free()
	arg2 := gobject.GBool(defaultEditable)
	defer arg2.Free()

	b := C.gtk_text_buffer_delete_selection(self.object, *((*C.gboolean)(arg1.GetPtr())), *((*C.gboolean)(arg2.GetPtr())))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextBuffer) PasteClipboard(clipboard *Clipboard, overrideLocation *TextIter, defaultEditable bool) {
	b := gobject.GBool(defaultEditable)
	defer b.Free()
	C.gtk_text_buffer_paste_clipboard(self.object, clipboard.object, &overrideLocation.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TextBuffer) CopyClipboard(clipboard *Clipboard) {
	C.gtk_text_buffer_copy_clipboard(self.object, clipboard.object)
}

func (self *TextBuffer) CutClipboard(clipboard *Clipboard, defaultEditable bool) {
	b := gobject.GBool(defaultEditable)
	defer b.Free()
	C.gtk_text_buffer_cut_clipboard(self.object, clipboard.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TextBuffer) GetSelectionBounds(start, end *TextIter) bool {
	b := C.gtk_text_buffer_get_selection_bounds(self.object, &start.object, &end.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextBuffer) BeginUserAction() {
	C.gtk_text_buffer_begin_user_action(self.object)
}

func (self *TextBuffer) EndUserAction() {
	C.gtk_text_buffer_end_user_action(self.object)
}

func (self *TextBuffer) AddSelectionClipboard(clipboard *Clipboard) {
	C.gtk_text_buffer_add_selection_clipboard(self.object, clipboard.object)
}

func (self *TextBuffer) RemoveSelectionClipboard(clipboard *Clipboard) {
	C.gtk_text_buffer_remove_selection_clipboard(self.object, clipboard.object)
}

//////////////////////////////
// End GtkTextBuffer
////////////////////////////// }}}

// GtkTextChildAnchor {{{
//////////////////////////////

// GtkTextChildAnchor type
type TextChildAnchor struct {
	object *C.GtkTextChildAnchor
}

func NewTextChildAnchor() *TextChildAnchor {
	t := &TextChildAnchor{}
	t.object = C.gtk_text_child_anchor_new()

	return t
}

// Conversion funcs
func newTextChildAnchorFromNative(obj unsafe.Pointer) interface{} {
	return &TextChildAnchor{(*C.GtkTextChildAnchor)(obj)}
}

func nativeFromTextChildAnchor(t interface{}) *gobject.GValue {
	tca, ok := t.(*TextChildAnchor)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TEXT_CHILD_ANCHOR, tca.ToNative())
		return gv
	}
	return nil
}

// TextChildAnchor is boxed type
func (self TextChildAnchor) GetTypeId() gobject.GType {
	return GtkType.TEXT_CHILD_ANCHOR
}

func (self TextChildAnchor) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

// TextChildAnchor interface
func (self *TextChildAnchor) GetWidgets() *glib.GSList {
	gl := C.gtk_text_child_anchor_get_widgets(self.object)

	goGSList := glib.NewGSListFromNative(unsafe.Pointer(gl))
	goGSList.GC_Free = true
	goGSList.GC_FreeFull = false
	glib.GSListFinalizer(goGSList)

	// Create conversion func
	goGSList.ConversionFunc = func(w unsafe.Pointer) interface{} {
		if widget, err := gobject.ConvertToGo(w); err == nil {
			return widget.(WidgetLike)
		}
		return nil
	}

	return goGSList
}

func (self *TextChildAnchor) GetDeleted() bool {
	b := C.gtk_text_child_anchor_get_deleted(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// End GtkTextChildAnchor
////////////////////////////// }}}

// GtkTextView {{{
//////////////////////////////

// TextView type
type TextView struct {
	object *C.GtkTextView
	*Container
}

// Create new TextView
func NewTextView() *TextView {
	t := &TextView{}
	o := C.gtk_text_view_new()
	t.object = C.to_GtkTextView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	}
	t.Container = NewContainer(unsafe.Pointer(o))
	textViewFinalizer(t)

	return t
}

func NewTextViewWithBuffer(buf *TextBuffer) *TextView {
	if buf == nil {
		return NewTextView()
	}
	t := &TextView{}
	o := C.gtk_text_view_new_with_buffer(buf.object)
	t.object = C.to_GtkTextView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	}
	t.Container = NewContainer(unsafe.Pointer(o))
	textViewFinalizer(t)

	return t
}

// Clear TextView struct when it goes out of reach
func textViewFinalizer(t *TextView) {
	runtime.SetFinalizer(t, func(t *TextView) { gobject.Unref(t) })
}

// Conversion function for gobject registration map
func newTextViewFromNative(obj unsafe.Pointer) interface{} {
	t := &TextView{}
	t.object = C.to_GtkTextView(obj)

	if gobject.IsObjectFloating(t) {
		gobject.RefSink(t)
	} else {
		gobject.Ref(t)
	}
	t.Container = NewContainer(obj)
	textViewFinalizer(t)

	return t
}

func nativeFromTextView(t interface{}) *gobject.GValue {
	text, ok := t.(*TextView)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TEXT_VIEW, text.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TextView) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TextView) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TextView) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TextView) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self TextView) C() *Container {
	return self.Container
}

// TextView interface

func (self *TextView) SetBuffer(buffer *TextBuffer) {
	C.gtk_text_view_set_buffer(self.object, buffer.object)
}

func (self *TextView) GetBuffer() *TextBuffer {
	b := C.gtk_text_view_get_buffer(self.object)

	if buf, err := gobject.ConvertToGo(unsafe.Pointer(b)); err == nil {
		return buf.(*TextBuffer)
	}
	return nil
}

func (self *TextView) ScrollToMark(mar *TextMark, withinMargin float64, useAlign bool, xalign float64, yalign float64) {
	b := gobject.GBool(useAlign)
	defer b.Free()
	C.gtk_text_view_scroll_to_mark(self.object, mar.object, C.gdouble(withinMargin), *((*C.gboolean)(b.GetPtr())), C.gdouble(xalign),
		C.gdouble(yalign))
}

func (self *TextView) ScrollToIter(iter *TextIter, withinMargin float64, useAlign bool, xalign, yalign float64) bool {
	b := gobject.GBool(useAlign)
	defer b.Free()
	res := C.gtk_text_view_scroll_to_iter(self.object, &iter.object, C.gdouble(withinMargin), *((*C.gboolean)(b.GetPtr())), C.gdouble(xalign),
		C.gdouble(yalign))
	return gobject.GoBool(unsafe.Pointer(&res))
}

func (self *TextView) ScrollMarkOnscreen(mark *TextMark) {
	C.gtk_text_view_scroll_mark_onscreen(self.object, mark.object)
}

func (self *TextView) MoveMarkOnscreen(mark *TextMark) {
	C.gtk_text_view_move_mark_onscreen(self.object, mark.object)
}

func (self *TextView) PlaceCursorOnscreen() bool {
	b := C.gtk_text_view_place_cursor_onscreen(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextView) GetVisibleRect() gdk3.Rectangle {
	var r C.GdkRectangle
	C.gtk_text_view_get_visible_rect(self.object, &r)

	if rec, err := gobject.ConvertToGo(unsafe.Pointer(&r), gdk3.GdkType.RECTANGLE); err == nil {
		return rec.(gdk3.Rectangle)
	}
	return gdk3.Rectangle{}
}

func (self *TextView) GetIterLocation(iter *TextIter) gdk3.Rectangle {
	var r C.GdkRectangle
	C.gtk_text_view_get_iter_location(self.object, &iter.object, &r)

	if rec, err := gobject.ConvertToGo(unsafe.Pointer(&r), gdk3.GdkType.RECTANGLE); err == nil {
		return rec.(gdk3.Rectangle)
	}
	return gdk3.Rectangle{}
}

func (self *TextView) GetCursorLocations(iter *TextIter) (strong, weak gdk3.Rectangle) {
	r1, r2 := new(C.GdkRectangle), new(C.GdkRectangle)
	C.gtk_text_view_get_cursor_locations(self.object, &iter.object, r1, r2)

	var rec1, rec2 gdk3.Rectangle
	if r1 != nil {
		if goRec1, err := gobject.ConvertToGo(unsafe.Pointer(r1), gdk3.GdkType.RECTANGLE); err == nil {
			rec1 = goRec1.(gdk3.Rectangle)
		}
	}

	if r2 != nil {
		if goRec2, err := gobject.ConvertToGo(unsafe.Pointer(r2), gdk3.GdkType.RECTANGLE); err == nil {
			rec2 = goRec2.(gdk3.Rectangle)
		}
	}
	return rec1, rec2
}

func (self *TextView) GetLineAtY(targetIter *TextIter, y int) (lineTop int) {
	var res C.gint
	C.gtk_text_view_get_line_at_y(self.object, &targetIter.object, C.gint(y), &res)
	lineTop = int(res)
	return
}

func (self *TextView) GetLineYrange(iter *TextIter) (y, height int) {
	var res, res1 C.gint
	C.gtk_text_view_get_line_yrange(self.object, &iter.object, &res, &res1)
	y = int(res)
	height = int(res1)
	return
}

func (self *TextView) GetIterAtLocation(iter *TextIter, x, y int) {
	C.gtk_text_view_get_iter_at_location(self.object, &iter.object, C.gint(x), C.gint(y))
}

func (self *TextView) GetIterAtPosition(iter *TextIter, x, y int) (trailing int) {
	var res C.gint
	C.gtk_text_view_get_iter_at_position(self.object, &iter.object, &res, C.gint(x), C.gint(y))
	return int(res)
}

func (self *TextView) BufferToWindowCoords(gtk_TextWindowType int, bufferX, bufferY int) (windowX, windowY int) {
	var res1, res2 C.gint
	C.gtk_text_view_buffer_to_window_coords(self.object, C.GtkTextWindowType(gtk_TextWindowType),
		C.gint(bufferX), C.gint(bufferY), &res1, &res2)

	windowX = int(res1)
	windowY = int(res2)
	return
}

func (self *TextView) WindowToBufferCoords(gtk_TextWindowType int, windowX, windowY int) (bufferX, bufferY int) {
	var res1, res2 C.gint
	C.gtk_text_view_window_to_buffer_coords(self.object, C.GtkTextWindowType(gtk_TextWindowType),
		C.gint(windowX), C.gint(windowY), &res1, &res2)

	bufferX = int(res1)
	bufferY = int(res2)
	return
}

func (self *TextView) GetWindow(gtk_TextWindowType int) *gdk3.Window {
	w := C.gtk_text_view_get_window(self.object, C.GtkTextWindowType(gtk_TextWindowType))

	if w == nil {
		return nil
	}

	if win, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
		return win.(*gdk3.Window)
	}
	return nil
}

func (self *TextView) GetWindowType(win *gdk3.Window) int {
	return int(C.gtk_text_view_get_window_type(self.object, (*C.GdkWindow)(win.ToNative())))
}

func (self *TextView) SetBorderWindowSize(gtk_TextWindowType, size int) {
	C.gtk_text_view_set_border_window_size(self.object, C.GtkTextWindowType(gtk_TextWindowType), C.gint(size))
}

func (self *TextView) GetBorderWindowSize(gtk_TextWindowType int) int {
	return int(C.gtk_text_view_get_border_window_size(self.object, C.GtkTextWindowType(gtk_TextWindowType)))
}

func (self *TextView) ForwardDisplayLine(iter *TextIter) bool {
	b := C.gtk_text_view_forward_display_line(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextView) BackwardDisplayLine(iter *TextIter) bool {
	b := C.gtk_text_view_backward_display_line(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextView) ForwardDisplayLineEnd(iter *TextIter) bool {
	b := C.gtk_text_view_forward_display_line_end(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextView) BackwardDisplayLineStart(iter *TextIter) bool {
	b := C.gtk_text_view_backward_display_line_start(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextView) StartsDisplayLine(iter *TextIter) bool {
	b := C.gtk_text_view_starts_display_line(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextView) MoveVisually(iter *TextIter, count int) bool {
	b := C.gtk_text_view_move_visually(self.object, &iter.object, C.gint(count))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextView) AddChildAtAnchor(child WidgetLike, anchor *TextChildAnchor) {
	C.gtk_text_view_add_child_at_anchor(self.object, child.W().object, anchor.object)
}

func (self *TextView) AddChildInWindow(child WidgetLike, gtk_TextWindowType, xpos, ypos int) {
	C.gtk_text_view_add_child_in_window(self.object, child.W().object, C.GtkTextWindowType(gtk_TextWindowType),
		C.gint(xpos), C.gint(ypos))
}

func (self *TextView) MoveChild(child WidgetLike, xpos, ypos int) {
	C.gtk_text_view_move_child(self.object, child.W().object, C.gint(xpos), C.gint(ypos))
}

func (self *TextView) SetWrapMode(gtk_WrapMode int) {
	C.gtk_text_view_set_wrap_mode(self.object, C.GtkWrapMode(gtk_WrapMode))
}

func (self *TextView) GetWrapMode() int {
	return int(C.gtk_text_view_get_wrap_mode(self.object))
}

func (self *TextView) SetEditable(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_text_view_set_editable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TextView) GetEditable() bool {
	b := C.gtk_text_view_get_editable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextView) SetCursorVisible(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_text_view_set_cursor_visible(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TextView) GetCursorVisible() bool {
	b := C.gtk_text_view_get_cursor_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextView) SetOverwrite(overwrite bool) {
	b := gobject.GBool(overwrite)
	defer b.Free()
	C.gtk_text_view_set_overwrite(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TextView) GetOverwrite() bool {
	b := C.gtk_text_view_get_overwrite(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TextView) SetPixelsAboveLines(pixelAbove int) {
	C.gtk_text_view_set_pixels_above_lines(self.object, C.gint(pixelAbove))
}

func (self *TextView) GetPixelsAboveLines() int {
	return int(C.gtk_text_view_get_pixels_above_lines(self.object))
}

func (self *TextView) SetPixelsBelowLines(pixelBelow int) {
	C.gtk_text_view_set_pixels_below_lines(self.object, C.gint(pixelBelow))
}

func (self *TextView) GetPixelsBelowLines() int {
	return int(C.gtk_text_view_get_pixels_below_lines(self.object))
}

func (self *TextView) SetPixelsInsideWrap(pixelInside int) {
	C.gtk_text_view_set_pixels_inside_wrap(self.object, C.gint(pixelInside))
}

func (self *TextView) GetPixelsInsideWrap() int {
	return int(C.gtk_text_view_get_pixels_inside_wrap(self.object))
}

func (self *TextView) SetJustification(gtk_Justification int) {
	C.gtk_text_view_set_justification(self.object, C.GtkJustification(gtk_Justification))
}

func (self *TextView) GetJustification() int {
	return int(C.gtk_text_view_get_justification(self.object))
}

func (self *TextView) SetLeftMargin(margin int) {
	C.gtk_text_view_set_left_margin(self.object, C.gint(margin))
}

func (self *TextView) GetLeftMargin() int {
	return int(C.gtk_text_view_get_left_margin(self.object))
}

func (self *TextView) SetRightMargin(margin int) {
	C.gtk_text_view_set_right_margin(self.object, C.gint(margin))
}

func (self *TextView) GetRightMargin() int {
	return int(C.gtk_text_view_get_right_margin(self.object))
}

func (self *TextView) SetIndent(indent int) {
	C.gtk_text_view_set_indent(self.object, C.gint(indent))
}

func (self *TextView) GetIndent() int {
	return int(C.gtk_text_view_get_indent(self.object))
}

//TODO: gtk_text_view_set_tabs
//TODO: gtk_text_view_get_tabs

func (self *TextView) SetAcceptsTab(acceptsTab bool) {
	b := gobject.GBool(acceptsTab)
	defer b.Free()
	C.gtk_text_view_set_accepts_tab(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TextView) GetAcceptsTab() bool {
	b := C.gtk_text_view_get_accepts_tab(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//TODO: gtk_text_view_get_default_attributes
//TODO: gtk_text_view_im_context_filter_keypress

func (self *TextView) ResetIMContext() {
	C.gtk_text_view_reset_im_context(self.object)
}

//////////////////////////////
// End GtkTextView
////////////////////////////// }}}

// End Multiline Text Editor }}}

// Tree, List and Icon Grid Widgets {{{

// GtkTreePath {{{
//////////////////////////////

// GtkTreePath type
type TreePath struct {
	object *C.GtkTreePath
}

func NewTreePath() *TreePath {
	tp := &TreePath{}
	tp.object = C.gtk_tree_path_new()
	treePathFinalizer(tp)

	return tp
}

func NewTreePathFromString(path string) *TreePath {
	tp := &TreePath{}

	s := gobject.GString(path)
	defer s.Free()

	tp.object = C.gtk_tree_path_new_from_string((*C.gchar)(s.GetPtr()))
	treePathFinalizer(tp)

	return tp
}

func NewTreePathFirst() *TreePath {
	tp := &TreePath{}
	tp.object = C.gtk_tree_path_new_first()
	treePathFinalizer(tp)

	return tp
}

// Clear TreePath struct when it goes out of scope
func treePathFinalizer(tp *TreePath) {
	runtime.SetFinalizer(tp, func(tp *TreePath) { tp.free() })
}

func newTreePathFromNative(obj unsafe.Pointer) interface{} {
	path := &TreePath{(*C.GtkTreePath)(obj)}
	return path
}

// TreePath Interface

//TODO: NewTreePathFromIndices

func (self *TreePath) free() {
	C.gtk_tree_path_free(self.object)
}

func (self *TreePath) ToString() string {
	s := C.gtk_tree_path_to_string(self.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *TreePath) AppendIndex(index int) {
	C.gtk_tree_path_append_index(self.object, C.gint(index))
}

func (self *TreePath) PrependIndex(index int) {
	C.gtk_tree_path_prepend_index(self.object, C.gint(index))
}

func (self *TreePath) GetDepth() int {
	return int(C.gtk_tree_path_get_depth(self.object))
}

func (self *TreePath) GetIndices() []int {
	cind := C.gtk_tree_path_get_indices(self.object)
	depth := self.GetDepth()
	gind := make([]int, depth)

	for i := 0; i < depth; i++ {
		indice := int(C._gtk_tree_path_get_indice(cind, C.int(i)))
		gind[i] = indice
	}

	return gind
}

func (self *TreePath) Copy() *TreePath {
	tp := &TreePath{}
	tp.object = C.gtk_tree_path_copy(self.object)
	treePathFinalizer(tp)

	return tp
}

func (self *TreePath) Compare(tp *TreePath) int {
	return int(C.gtk_tree_path_compare(self.object, tp.object))
}

func (self *TreePath) Next() {
	C.gtk_tree_path_next(self.object)
}

func (self *TreePath) Prev() bool {
	b := C.gtk_tree_path_prev(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreePath) Up() bool {
	b := C.gtk_tree_path_up(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreePath) Down() {
	C.gtk_tree_path_down(self.object)
}

func (self *TreePath) IsAncestor(descendant *TreePath) bool {
	b := C.gtk_tree_path_is_ancestor(self.object, descendant.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreePath) IsDescendant(ancestor *TreePath) bool {
	b := C.gtk_tree_path_is_descendant(self.object, ancestor.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// END GtkTreePath
////////////////////////////// }}}

// GtkTreeRowReference {{{
//////////////////////////////

// GtkTreeRowReference type
type TreeRowReference struct {
	object *C.GtkTreeRowReference
}

//func NewTreeRowReference() *TreeRowReference {

// Clear TreeRowReference when it goes out of reach
func treeRowReferenceFinalizer(trr *TreeRowReference) {
	runtime.SetFinalizer(trr, func(trr *TreeRowReference) { trr.free() })
}

// TreeRowReference interface
func (self *TreeRowReference) free() {
	C.gtk_tree_row_reference_free(self.object)
}

//////////////////////////////
// END TreeRowReference
////////////////////////////// }}}

// GtkTreeIter {{{
//////////////////////////////

// GtkTreeRowReference type
type TreeIter struct {
	object C.GtkTreeIter
}

func (self *TreeIter) Copy() *TreeIter {
	ti := C.gtk_tree_iter_copy(&self.object)
	gti := &TreeIter{*ti}
	C.gtk_tree_iter_free(ti)

	return gti
}

//////////////////////////////
// END GtkTreeIter
////////////////////////////// }}}

// GtkTreeModel {{{
//////////////////////////////

// GtkTreeModel type
type TreeModel struct {
	object *C.GtkTreeModel
}

func NewTreeModel(ctm unsafe.Pointer) *TreeModel {
	tm := &TreeModel{}
	tm.object = C.to_GtkTreeModel(ctm)

	if gobject.IsObjectFloating(tm) {
		gobject.RefSink(tm)
	} else {
		gobject.Ref(tm)
	}
	treeModelFinalizer(tm)

	return tm
}

// Clear TreeModel struct when it goes out of scope
func treeModelFinalizer(tm *TreeModel) {
	runtime.SetFinalizer(tm, func(tm *TreeModel) { gobject.Unref(tm) })
}

// Conversion funcs
func newTreeModelFromNative(obj unsafe.Pointer) interface{} {
	return NewTreeModel(obj)
}

func nativeFromTreeModel(tm interface{}) *gobject.GValue {
	tree, ok := tm.(*TreeModel)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TREE_MODEL, tree.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TreeModel) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TreeModel) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TreeModel) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TreeModel) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// TreeModel interface

func (self *TreeModel) GetFlags() int {
	return int(C.gtk_tree_model_get_flags(self.object))
}

func (self *TreeModel) GetNColumns() int {
	return int(C.gtk_tree_model_get_n_columns(self.object))
}

func (self *TreeModel) GetColumnType(index int) gobject.GType {
	return gobject.GType(C.gtk_tree_model_get_column_type(self.object, C.gint(index)))
}

func (self *TreeModel) GetIter(iter *TreeIter, path *TreePath) bool {
	b := C.gtk_tree_model_get_iter(self.object, &iter.object, path.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) GetIterFromString(iter *TreeIter, pathString string) bool {
	s := gobject.GString(pathString)
	defer s.Free()
	b := C.gtk_tree_model_get_iter_from_string(self.object, &iter.object, (*C.gchar)(s.GetPtr()))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) GetIterFirst(iter *TreeIter) bool {
	b := C.gtk_tree_model_get_iter_first(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) GetPath(iter *TreeIter) *TreePath {
	p := C.gtk_tree_model_get_path(self.object, &iter.object)
	path := &TreePath{p}
	treePathFinalizer(path)

	return path
}

func (self *TreeModel) GetValue(iter *TreeIter, column int) interface{} {
	var val C.GValue
	C.gtk_tree_model_get_value(self.object, &iter.object, C.gint(column), &val)

	gv := gobject.NewGValueFromNative(unsafe.Pointer(&val))
	defer gv.Free()
	t, e := gobject.ConvertToGo(gv.GetPtr(), gv.GetTypeID())

	if e == nil {
		return t
	}
	return nil
}

func (self *TreeModel) IterNext(iter *TreeIter) bool {
	b := C.gtk_tree_model_iter_next(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) IterPrevious(iter *TreeIter) bool {
	b := C.gtk_tree_model_iter_previous(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) IterChildren(iter *TreeIter, parent *TreeIter) bool {
	b := C.gtk_tree_model_iter_children(self.object, &iter.object, &parent.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) IterHasChild(iter *TreeIter) bool {
	b := C.gtk_tree_model_iter_has_child(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) IterNChildren(iter *TreeIter) int {
	return int(C.gtk_tree_model_iter_n_children(self.object, &iter.object))
}

func (self *TreeModel) IterNthChild(iter *TreeIter, parent *TreeIter, n int) bool {
	b := C.gtk_tree_model_iter_nth_child(self.object, &iter.object, &parent.object, C.gint(n))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) IterParent(iter *TreeIter, child *TreeIter) bool {
	b := C.gtk_tree_model_iter_parent(self.object, &iter.object, &child.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) GetStringFromIter(iter *TreeIter) string {
	s := C.gtk_tree_model_get_string_from_iter(self.object, &iter.object)
	return gobject.GoString(unsafe.Pointer(&s))
}

//TODO: gtk_tree_model_ref_node
//TODO: gtk_tree_model_unref_node
//TODO: gtk_tree_model_get_valist
//TODO: gtk_tree_model_foreach
//...
//////////////////////////////
// END GtkTreeModel
////////////////////////////// }}}

// GtkListStore {{{
//////////////////////////////

// GtkListStore type
type ListStore struct {
	object *C.GtkListStore
	*TreeModel
}

func NewListStore(gtypeSlice []gobject.GType) *ListStore {
	count := len(gtypeSlice)
	if count == 0 {
		return nil
	}

	arr := C.gtype_array_new(C.int(count))
	defer C.gtype_array_free(arr)

	for i, gtype := range gtypeSlice {
		C.gtype_array_set_element(arr, C.int(i), C.GType(gtype))
	}

	ls := &ListStore{}
	ls.object = C.gtk_list_store_newv(C.gint(count), arr)

	if gobject.IsObjectFloating(ls) {
		gobject.RefSink(ls)
	}
	ls.TreeModel = NewTreeModel(unsafe.Pointer(ls.object))
	listStoreFinalizer(ls)

	return ls
}

// Clear ListStore struct when it goes out of reach
func listStoreFinalizer(ls *ListStore) {
	runtime.SetFinalizer(ls, func(ls *ListStore) { gobject.Unref(ls) })
}

// Conversion functions
func newListStoreFromNative(obj unsafe.Pointer) interface{} {
	ls := &ListStore{}
	ls.object = C.to_GtkListStore(obj)

	if gobject.IsObjectFloating(ls) {
		gobject.RefSink(ls)
	} else {
		gobject.Ref(ls)
	}
	ls.TreeModel = NewTreeModel(obj)
	listStoreFinalizer(ls)

	return ls
}

func nativeFromListStore(ls interface{}) *gobject.GValue {
	listStore, ok := ls.(*ListStore)
	if ok {
		gv := gobject.CreateCGValue(GtkType.LIST_STORE, listStore.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self ListStore) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ListStore) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ListStore) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ListStore) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// ListStore implements TreeModel
func (self ListStore) ITreeModel() *TreeModel {
	return self.TreeModel
}

// ListStore interface

func (self *ListStore) SetValue(iter *TreeIter, column int, value interface{}) {
	var cval *gobject.GValue
	if value == nil {
		cval = gobject.CreateCGValue(self.GetColumnType(column))
	} else {
		cval = gobject.ConvertToC(value)

		// If value cannot be converted to fundamental
		// gobject value, then we don't do anything
		if cval == nil {
			return
		}
	}
	defer cval.Free()

	C.gtk_list_store_set_value(self.object, &iter.object, C.gint(column), (*C.GValue)(cval.ToCGValue()))
}

func (self *ListStore) SetValues(iter *TreeIter, values map[int]interface{}) {
	for k, v := range values {
		self.SetValue(iter, k, v)
	}
}

func (self *ListStore) Remove(iter *TreeIter) bool {
	b := C.gtk_list_store_remove(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ListStore) Insert(iter *TreeIter, position int) {
	C.gtk_list_store_insert(self.object, &iter.object, C.gint(position))
}

func (self *ListStore) InsertBefore(iter, sibling *TreeIter) {
	var cSibling *C.GtkTreeIter = nil

	if sibling != nil {
		cSibling = &sibling.object
	}
	C.gtk_list_store_insert_before(self.object, &iter.object, cSibling)
}

func (self *ListStore) InsertAfter(iter, sibling *TreeIter) {
	var cSibling *C.GtkTreeIter = nil

	if sibling != nil {
		cSibling = &sibling.object
	}
	C.gtk_list_store_insert_after(self.object, &iter.object, cSibling)
}

func (self *ListStore) InsertWithValues(iter *TreeIter, position int, values map[int]interface{}) {
	if values == nil || len(values) == 0 {
		return
	}

	valArr := C.gvalue_array_new(C.int(len(values)))
	columns := C.gint_array_new(C.int(len(values)))
	defer C.gvalue_array_free(valArr)
	defer C.gint_array_free(columns)

	var i int = 0
	for k, v := range values {
		C.gint_array_set_element(columns, C.int(i), C.gint(k))
		cval := gobject.ConvertToC(v)
		defer cval.Free()
		C.gvalue_array_set_element(valArr, C.int(i), (*C.GValue)(cval.ToCGValue()))
		i++
	}
	var Citer *C.GtkTreeIter = nil

	if iter != nil {
		Citer = &iter.object
	}

	C.gtk_list_store_insert_with_valuesv(self.object, Citer, C.gint(position), columns, valArr, C.gint(len(values)))
}

func (self *ListStore) Prepend(iter *TreeIter) {
	C.gtk_list_store_prepend(self.object, &iter.object)
}

func (self *ListStore) Append(iter *TreeIter) {
	C.gtk_list_store_append(self.object, &iter.object)
}

func (self *ListStore) Clear() {
	C.gtk_list_store_clear(self.object)
}

//TODO: gtk_list_store_reorder

func (self *ListStore) Swap(a, b *TreeIter) {
	C.gtk_list_store_swap(self.object, &a.object, &b.object)
}

func (self *ListStore) MoveBefore(iter, position *TreeIter) {
	C.gtk_list_store_move_before(self.object, &iter.object, &position.object)
}

func (self *ListStore) MoveAfter(iter, position *TreeIter) {
	C.gtk_list_store_move_after(self.object, &iter.object, &position.object)
}

//////////////////////////////
// END GtkListStore
////////////////////////////// }}}

// GtkTreeStore {{{
//////////////////////////////

// GtkTreeStore type
type TreeStore struct {
	object *C.GtkTreeStore
	*TreeModel
}

func NewTreeStore(gtypeSlice []gobject.GType) *TreeStore {
	count := len(gtypeSlice)
	if count == 0 {
		return nil
	}

	arr := C.gtype_array_new(C.int(count))
	defer C.gtype_array_free(arr)

	for i, gtype := range gtypeSlice {
		C.gtype_array_set_element(arr, C.int(i), C.GType(gtype))
	}

	ts := &TreeStore{}
	ts.object = C.gtk_tree_store_newv(C.gint(count), arr)

	if gobject.IsObjectFloating(ts) {
		gobject.RefSink(ts)
	}
	ts.TreeModel = NewTreeModel(unsafe.Pointer(ts.object))
	treeStoreFinalizer(ts)

	return ts
}

// Clear TreeStore struct when it goes out of reach
func treeStoreFinalizer(ts *TreeStore) {
	runtime.SetFinalizer(ts, func(ts *TreeStore) { gobject.Unref(ts) })
}

// Conversion functions
func newTreeStoreFromNative(obj unsafe.Pointer) interface{} {
	ts := &TreeStore{}
	ts.object = C.to_GtkTreeStore(obj)

	if gobject.IsObjectFloating(ts) {
		gobject.RefSink(ts)
	} else {
		gobject.Ref(ts)
	}
	ts.TreeModel = NewTreeModel(obj)
	treeStoreFinalizer(ts)

	return ts
}

func nativeFromTreeStore(ts interface{}) *gobject.GValue {
	treeStore, ok := ts.(*TreeStore)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TREE_STORE, treeStore.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TreeStore) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TreeStore) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TreeStore) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TreeStore) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// TreeStore implements TreeModel
func (self TreeStore) ITreeModel() *TreeModel {
	return self.TreeModel
}

// TreeStore interface

func (self *TreeStore) SetValue(iter *TreeIter, column int, value interface{}) {
	var cval *gobject.GValue
	if value == nil {
		cval = gobject.CreateCGValue(self.GetColumnType(column))
	} else {
		cval = gobject.ConvertToC(value)

		// If value cannot be converted to fundamnetal
		// gobject value, then we don't do anything
		if cval == nil {
			return
		}
	}
	defer cval.Free()

	C.gtk_tree_store_set_value(self.object, &iter.object, C.gint(column), (*C.GValue)(cval.ToCGValue()))
}

func (self *TreeStore) SetValues(iter *TreeIter, values map[int]interface{}) {
	for k, v := range values {
		self.SetValue(iter, k, v)
	}
}

func (self *TreeStore) Remove(iter *TreeIter) bool {
	b := C.gtk_tree_store_remove(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeStore) Insert(iter, parent *TreeIter, position int) {
	var cparent *C.GtkTreeIter = nil

	if parent != nil {
		cparent = &parent.object
	}
	C.gtk_tree_store_insert(self.object, &iter.object, cparent, C.gint(position))
}

func (self *TreeStore) InsertBefore(iter, parent, sibling *TreeIter) {
	var cSibling, cParent *C.GtkTreeIter = nil, nil

	if sibling != nil {
		cSibling = &sibling.object
	}

	if parent != nil {
		cParent = &parent.object
	}
	C.gtk_tree_store_insert_before(self.object, &iter.object, cParent, cSibling)
}

func (self *TreeStore) InsertAfter(iter, parent, sibling *TreeIter) {
	var cSibling, cParent *C.GtkTreeIter = nil, nil

	if sibling != nil {
		cSibling = &sibling.object
	}

	if parent != nil {
		cParent = &parent.object
	}
	C.gtk_tree_store_insert_after(self.object, &iter.object, cParent, cSibling)
}

func (self *TreeStore) InsertWithValues(iter, parent *TreeIter, position int, values map[int]interface{}) {
	if values == nil || len(values) == 0 {
		return
	}

	valArr := C.gvalue_array_new(C.int(len(values)))
	columns := C.gint_array_new(C.int(len(values)))
	defer C.gvalue_array_free(valArr)
	defer C.gint_array_free(columns)

	var i int = 0
	for k, v := range values {
		C.gint_array_set_element(columns, C.int(i), C.gint(k))
		cval := gobject.ConvertToC(v)
		defer cval.Free()
		C.gvalue_array_set_element(valArr, C.int(i), (*C.GValue)(cval.ToCGValue()))
		i++
	}
	var Citer, Cparent *C.GtkTreeIter = nil, nil

	if iter != nil {
		Citer = &iter.object
	}

	if parent != nil {
		Cparent = &parent.object
	}

	C.gtk_tree_store_insert_with_valuesv(self.object, Citer, Cparent, C.gint(position), columns, valArr, C.gint(len(values)))
}

func (self *TreeStore) Prepend(iter, parent *TreeIter) {
	var cParent *C.GtkTreeIter = nil

	if parent != nil {
		cParent = &parent.object
	}
	C.gtk_tree_store_prepend(self.object, &iter.object, cParent)
}

func (self *TreeStore) Append(iter, parent *TreeIter) {
	var cParent *C.GtkTreeIter = nil
	if parent != nil {
		cParent = &parent.object
	}
	C.gtk_tree_store_append(self.object, &iter.object, cParent)
}

func (self *TreeStore) IsAncestor(iter, descendant *TreeIter) bool {
	b := C.gtk_tree_store_is_ancestor(self.object, &iter.object, &descendant.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeStore) IterDepth(iter *TreeIter) int {
	return int(C.gtk_tree_store_iter_depth(self.object, &iter.object))
}

func (self *TreeStore) Clear() {
	C.gtk_tree_store_clear(self.object)
}

//TODO: gtk_tree_store_reorder

func (self *TreeStore) Swap(a, b *TreeIter) {
	C.gtk_tree_store_swap(self.object, &a.object, &b.object)
}

func (self *TreeStore) MoveBefore(iter, position *TreeIter) {
	C.gtk_tree_store_move_before(self.object, &iter.object, &position.object)
}

func (self *TreeStore) MoveAfter(iter, position *TreeIter) {
	C.gtk_tree_store_move_after(self.object, &iter.object, &position.object)
}

//////////////////////////////
// END GtkTreeStore
////////////////////////////// }}}

// GtkCellArea {{{
//////////////////////////////

// GtkCellArea type
type CellArea struct {
	object *C.GtkCellArea
}

// Clear CellArea struct when it goes out of reach
func cellAreaFinalizer(ca *CellArea) {
	runtime.SetFinalizer(ca, func(ca *CellArea) { gobject.Unref(ca) })
}

// Conversion funcs
func newCellAreaFromNative(obj unsafe.Pointer) interface{} {
	ca := &CellArea{}
	ca.object = C.to_GtkCellArea(obj)

	if gobject.IsObjectFloating(ca) {
		gobject.RefSink(ca)
	} else {
		gobject.Ref(ca)
	}
	cellAreaFinalizer(ca)

	return ca
}

func nativeFromCellArea(ca interface{}) *gobject.GValue {
	cell, ok := ca.(*CellArea)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_AREA, cell.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellArea) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellArea) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellArea) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellArea) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// CellArea interface

func (self *CellArea) Add(renderer CellRendererLike) {
	C.gtk_cell_area_add(self.object, renderer.CRenderer().object)
}

func (self *CellArea) Remove(renderer CellRendererLike) {
	C.gtk_cell_area_remove(self.object, renderer.CRenderer().object)
}

func (self *CellArea) HasRenderer(renderer CellRendererLike) bool {
	b := C.gtk_cell_area_has_renderer(self.object, renderer.CRenderer().object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellArea) Foreach(callback interface{}, data ...interface{}) {
	f, id := gobject.CreateCustomClosure(callback, data...)
	_closures[id] = f
	C._gtk_cell_area_foreach(self.object, C.gint64(id))
	delete(_closures, id)
}

//////////////////////////////
// End GtkCellArea
////////////////////////////// }}}

// GtkCellAreaContext {{{
//////////////////////////////

// GtkCellAreaContext type
type CellAreaContext struct {
	object *C.GtkCellAreaContext
}

// Clear CellAreaContext struct when it goes out of reach
func cellAreaContextFinalizer(ca *CellAreaContext) {
	runtime.SetFinalizer(ca, func(ca *CellAreaContext) { gobject.Unref(ca) })
}

// Conversion funcs
func newCellAreaContextFromNative(obj unsafe.Pointer) interface{} {
	ca := &CellAreaContext{}
	ca.object = C.to_GtkCellAreaContext(obj)

	if gobject.IsObjectFloating(ca) {
		gobject.RefSink(ca)
	} else {
		gobject.Ref(ca)
	}
	cellAreaContextFinalizer(ca)

	return ca
}

func nativeFromCellAreaContext(ca interface{}) *gobject.GValue {
	cell, ok := ca.(*CellAreaContext)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_AREA_CONTEXT, cell.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellAreaContext) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellAreaContext) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellAreaContext) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellAreaContext) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// CellAreaContext interface

func (self *CellAreaContext) GetArea() *CellArea {
	ar := C.gtk_cell_area_context_get_area(self.object)
	if ar == nil {
		return nil
	}
	area, _ := gobject.ConvertToGo(unsafe.Pointer(ar))

	return area.(*CellArea)
}

func (self *CellAreaContext) Allocate(width, height int) {
	C.gtk_cell_area_context_allocate(self.object, C.gint(width), C.gint(height))
}

func (self *CellAreaContext) Reset() {
	C.gtk_cell_area_context_reset(self.object)
}

func (self *CellAreaContext) GetPreferredWidth() (mWidth, nWidth int) {
	var m, n C.gint
	C.gtk_cell_area_context_get_preferred_width(self.object, &m, &n)
	mWidth = int(m)
	nWidth = int(n)
	return
}

func (self *CellAreaContext) GetPreferredHeight() (mHeight, nHeight int) {
	var m, n C.gint
	C.gtk_cell_area_context_get_preferred_width(self.object, &m, &n)
	mHeight = int(m)
	nHeight = int(n)
	return
}

func (self *CellAreaContext) GetPreferredHeightForWidth(width int) (mHeight, nHeight int) {
	var m, n C.gint
	C.gtk_cell_area_context_get_preferred_height_for_width(self.object, C.gint(width), &m, &n)
	mHeight = int(m)
	nHeight = int(n)
	return
}

func (self *CellAreaContext) GetPreferredWidthForHeight(height int) (mWidth, nWidth int) {
	var m, n C.gint
	C.gtk_cell_area_context_get_preferred_width_for_height(self.object, C.gint(height), &m, &n)
	mWidth = int(m)
	nWidth = int(n)
	return
}

func (self *CellAreaContext) GetAllocation() (width, height int) {
	var w, h C.gint
	C.gtk_cell_area_context_get_allocation(self.object, &w, &h)
	width = int(w)
	height = int(h)
	return
}

func (self *CellAreaContext) PushPreferredWidth(minimumWidth, naturalWidth int) {
	C.gtk_cell_area_context_push_preferred_width(self.object, C.gint(minimumWidth), C.gint(naturalWidth))
}

func (self *CellAreaContext) PushPreferredHeight(minimumHeight, naturalHeight int) {
	C.gtk_cell_area_context_push_preferred_height(self.object, C.gint(minimumHeight), C.gint(naturalHeight))
}

//////////////////////////////
// End GtkCellAreaContext
////////////////////////////// }}}

// GtkCellLayout interface {{{
//////////////////////////////

// GtkCellLayout type
type CellLayout struct {
	object *C.GtkCellLayout
}

// Clear CellLayout when it goes out of reach
func cellLayoutFinalizer(cl *CellLayout) {
	runtime.SetFinalizer(cl, func(cl *CellLayout) { gobject.Unref(cl) })
}

// Conversion funcs
func newCellLayoutFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellLayout{}
	cl.object = C.to_GtkCellLayout(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cellLayoutFinalizer(cl)

	return cl
}

func nativeFromCellLayout(cl interface{}) *gobject.GValue {
	cell, ok := cl.(*CellLayout)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_LAYOUT, cell.ToNative())
		return gv
	}
	return nil
}

// Simulate object-like
func (self CellLayout) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellLayout) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellLayout) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellLayout) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// CellLayout interface

func (self *CellLayout) PackStart(cell CellRendererLike, expand bool) {
	b := gobject.GBool(expand)
	defer b.Free()
	C.gtk_cell_layout_pack_start(self.object, cell.CRenderer().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellLayout) PackEnd(cell CellRendererLike, expand bool) {
	b := gobject.GBool(expand)
	defer b.Free()
	C.gtk_cell_layout_pack_end(self.object, cell.CRenderer().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellLayout) GetArea() *CellArea {
	ar := C.gtk_cell_layout_get_area(self.object)
	if ar == nil {
		return nil
	}

	area, err := gobject.ConvertToGo(unsafe.Pointer(ar))
	if err == nil {
		return area.(*CellArea)
	}
	return nil
}

func (self *CellLayout) GetCells() *glib.GList {
	gl := C.gtk_cell_layout_get_cells(self.object)
	if gl == nil {
		return nil
	}

	goGlist := glib.NewGListFromNative(unsafe.Pointer(gl))

	// this GList must be freed, but not cell renderers inside
	goGlist.GC_Free = true
	goGlist.GC_FreeFull = false
	// Set finalizer
	glib.GListFinalizer(goGlist)

	// CellRenderers are inside this GList so we add conversion func
	// that converts C CellRenderers to Go CellRenderer struct
	goGlist.ConversionFunc = func(obj unsafe.Pointer) interface{} {
		renderer, err := gobject.ConvertToGo(obj)
		if err == nil {
			return renderer.(CellRendererLike)
		}
		return nil
	}

	return goGlist
}

func (self *CellLayout) Reorder(cell CellRendererLike, position int) {
	C.gtk_cell_layout_reorder(self.object, cell.CRenderer().object, C.gint(position))
}

func (self *CellLayout) Clear() {
	C.gtk_cell_layout_clear(self.object)
}

func (self *CellLayout) AddAttribute(cell CellRendererLike, attribute string, column int) {
	s := gobject.GString(attribute)
	defer s.Free()
	C.gtk_cell_layout_add_attribute(self.object, cell.CRenderer().object, (*C.gchar)(s.GetPtr()), C.gint(column))
}

func (self *CellLayout) ClearAttributes(cell CellRendererLike) {
	C.gtk_cell_layout_clear_attributes(self.object, cell.CRenderer().object)
}

func (self *CellLayout) SetAttributes(cell CellRendererLike, attributes A) {
	self.ClearAttributes(cell)

	for _, attr := range attributes {
		self.AddAttribute(cell, attr.Attribute, attr.Column)
	}
}

func (self *CellLayout) SetCellDataFunc(renderer CellRendererLike, f interface{}, data ...interface{}) {
	cl, id := gobject.CreateCustomClosure(f, data...)
	_closures[id] = cl
	C._gtk_cell_layout_set_cell_data_func(self.object, renderer.CRenderer().object, C.gint64(id))
}

//////////////////////////////
// End GtkCellLayout interface
////////////////////////////// }}}

// GtkCellView {{{
//////////////////////////////

// GtkCellView type
type CellView struct {
	object *C.GtkCellView
	*Widget
	*Orientable
}

func NewCellView() *CellView {
	cv := &CellView{}
	o := C.gtk_cell_view_new()
	cv.object = C.to_GtkCellView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cv) {
		gobject.RefSink(cv)
	}
	cv.Widget = NewWidget(unsafe.Pointer(o))
	cv.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	cellViewFinalizer(cv)

	return cv
}

func NewCellViewWithContext(area CellAreaLike, context *CellAreaContext) *CellView {
	cv := &CellView{}
	o := C.gtk_cell_view_new_with_context(area.CArea().object, context.object)
	cv.object = C.to_GtkCellView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cv) {
		gobject.RefSink(cv)
	}
	cv.Widget = NewWidget(unsafe.Pointer(o))
	cv.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	cellViewFinalizer(cv)

	return cv
}

func NewCellViewWithText(text string) *CellView {
	cv := &CellView{}
	s := gobject.GString(text)
	defer s.Free()

	o := C.gtk_cell_view_new_with_text((*C.gchar)(s.GetPtr()))
	cv.object = C.to_GtkCellView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cv) {
		gobject.RefSink(cv)
	}
	cv.Widget = NewWidget(unsafe.Pointer(o))
	cv.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	cellViewFinalizer(cv)

	return cv
}

func NewCellViewWithMarkup(markup string) *CellView {
	cv := &CellView{}
	s := gobject.GString(markup)
	defer s.Free()

	o := C.gtk_cell_view_new_with_markup((*C.gchar)(s.GetPtr()))
	cv.object = C.to_GtkCellView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cv) {
		gobject.RefSink(cv)
	}
	cv.Widget = NewWidget(unsafe.Pointer(o))
	cv.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	cellViewFinalizer(cv)

	return cv
}

func NewCellViewWithPixbuf(pixbuf *gdkpixbuf.Pixbuf) *CellView {
	cv := &CellView{}
	o := C.gtk_cell_view_new_with_pixbuf((*C.GdkPixbuf)(pixbuf.ToNative()))
	cv.object = C.to_GtkCellView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cv) {
		gobject.RefSink(cv)
	}
	cv.Widget = NewWidget(unsafe.Pointer(o))
	cv.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	cellViewFinalizer(cv)

	return cv
}

// Clear CellView struct when it goes out of reach
func cellViewFinalizer(cv *CellView) {
	runtime.SetFinalizer(cv, func(cv *CellView) { gobject.Unref(cv) })
}

// Conversion funcs
func newCellViewFromNative(obj unsafe.Pointer) interface{} {
	cv := &CellView{}
	cv.object = C.to_GtkCellView(obj)

	if gobject.IsObjectFloating(cv) {
		gobject.RefSink(cv)
	} else {
		gobject.Ref(cv)
	}
	cv.Widget = NewWidget(obj)
	cv.Orientable = newOrientableFromNative(obj).(*Orientable)
	cellViewFinalizer(cv)

	return cv
}

func nativeFromCellView(cv interface{}) *gobject.GValue {
	cell, ok := cv.(*CellView)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_VIEW, cell.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellView) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellView) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellView) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellView) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self CellView) W() *Widget {
	return self.Widget
}

// To be Orientable-like
func (self CellView) IOrientable() *Orientable {
	return self.Orientable
}

// CellView interface

func (self *CellView) SetModel(model TreeModelLike) {
	if model == nil {
		C.gtk_cell_view_set_model(self.object, nil)
	}
	C.gtk_cell_view_set_model(self.object, model.ITreeModel().object)
}

func (self *CellView) GetModel() TreeModelLike {
	model := C.gtk_cell_view_get_model(self.object)
	if model == nil {
		return nil
	}

	if m, err := gobject.ConvertToGo(unsafe.Pointer(model)); err == nil {
		return m.(TreeModelLike)
	}
	return nil
}

func (self *CellView) SetDisplayedRow(path *TreePath) {
	C.gtk_cell_view_set_displayed_row(self.object, path.object)
}

func (self *CellView) GetDisplayedRow() *TreePath {
	p := C.gtk_cell_view_get_displayed_row(self.object)

	if p == nil {
		return nil
	}

	if path, err := gobject.ConvertToGo(unsafe.Pointer(p)); err == nil {
		return path.(*TreePath)
	}
	return nil
}

func (self *CellView) SetBackgrundColor(color gdk3.Color) {
	c := gobject.ConvertToC(color)
	defer c.Free()
	C.gtk_cell_view_set_background_color(self.object, (*C.GdkColor)(c.GetPtr()))
}

func (self *CellView) SetBackgroundRGBA(rgba gdk3.RGBA) {
	r := gobject.ConvertToC(rgba)
	defer r.Free()
	C.gtk_cell_view_set_background_rgba(self.object, (*C.GdkRGBA)(r.GetPtr()))
}

func (self *CellView) SetDrawSensitive(drawSensitive bool) {
	b := gobject.GBool(drawSensitive)
	defer b.Free()
	C.gtk_cell_view_set_draw_sensitive(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellView) GetDrawSensitive() bool {
	b := C.gtk_cell_view_get_draw_sensitive(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellView) SetFitModel(fitModel bool) {
	b := gobject.GBool(fitModel)
	defer b.Free()
	C.gtk_cell_view_set_fit_model(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellView) GetFitModel() bool {
	b := C.gtk_cell_view_get_fit_model(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// End GtkCellView
////////////////////////////// }}}

// GtkCellRenderer {{{
//////////////////////////////

// GtkCellRenderer type
type CellRenderer struct {
	object *C.GtkCellRenderer
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererFinalizer(cl *CellRenderer) {
	runtime.SetFinalizer(cl, func(cl *CellRenderer) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRenderer{}
	cl.object = C.to_GtkCellRenderer(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cellRendererFinalizer(cl)

	return cl
}

func nativeFromCellRenderer(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRenderer)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRenderer) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRenderer) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRenderer) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRenderer) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Interface functions

func (self *CellRenderer) GetAlignedArea(w WidgetLike, gtk_CellRendererState_flags int, cellArea gdk3.Rectangle) gdk3.Rectangle {
	ca := gobject.ConvertToC(cellArea)
	defer ca.Free()

	var alignedArea C.GdkRectangle

	C.gtk_cell_renderer_get_aligned_area(self.object, w.W().object, C.GtkCellRendererState(gtk_CellRendererState_flags),
		(*C.GdkRectangle)(ca.GetPtr()), &alignedArea)

	if rec, err := gobject.ConvertToGo(unsafe.Pointer(&alignedArea), gdk3.GdkType.RECTANGLE); err == nil {
		return rec.(gdk3.Rectangle)
	}
	return gdk3.Rectangle{}
}

//TODO: gtk_cell_renderer_render
//TODO: gtk_cell_renderer_activate
//TODO: gtk_cell_renderer_start_editing
//TODO: gtk_cell_renderer_stop_editing

func (self *CellRenderer) GetFixedSize() (width, height int) {
	var w C.gint
	var h C.gint
	C.gtk_cell_renderer_get_fixed_size(self.object, &w, &h)

	width = int(w)
	height = int(h)
	return
}

func (self *CellRenderer) SetFixedSize(width, height int) {
	C.gtk_cell_renderer_set_fixed_size(self.object, C.gint(width), C.gint(height))
}

func (self *CellRenderer) GetVisible() bool {
	b := C.gtk_cell_renderer_get_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRenderer) SetVisible(visible bool) {
	b := gobject.GBool(visible)
	defer b.Free()
	C.gtk_cell_renderer_set_visible(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellRenderer) GetSensitive() bool {
	b := C.gtk_cell_renderer_get_sensitive(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRenderer) SetSensitive(sensitive bool) {
	b := gobject.GBool(sensitive)
	defer b.Free()
	C.gtk_cell_renderer_set_sensitive(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellRenderer) GetAlignment() (xalign, yalign float32) {
	var x C.gfloat
	var y C.gfloat
	C.gtk_cell_renderer_get_alignment(self.object, &x, &y)

	xalign = float32(x)
	yalign = float32(y)
	return
}

func (self *CellRenderer) SetAlignment(xalign, yalign float32) {
	C.gtk_cell_renderer_set_alignment(self.object, C.gfloat(xalign), C.gfloat(yalign))
}

func (self *CellRenderer) GetPadding() (xpad, ypad int) {
	var x C.gint
	var y C.gint
	C.gtk_cell_renderer_get_padding(self.object, &x, &y)
	xpad = int(x)
	ypad = int(y)
	return
}

func (self *CellRenderer) SetPadding(xpad, ypad int) {
	C.gtk_cell_renderer_set_padding(self.object, C.gint(xpad), C.gint(ypad))
}

func (self *CellRenderer) GetState(w WidgetLike, gtk_CellRendererState int) int {
	return int(C.gtk_cell_renderer_get_state(self.object, w.W().object, C.GtkCellRendererState(gtk_CellRendererState)))
}

func (self *CellRenderer) IsActivatable() bool {
	b := C.gtk_cell_renderer_is_activatable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRenderer) GetPreferredHeight(w WidgetLike) (minimumSize, naturalSize int) {
	var m, n C.gint
	C.gtk_cell_renderer_get_preferred_height(self.object, w.W().object, &m, &n)
	minimumSize = int(m)
	naturalSize = int(n)
	return
}

func (self *CellRenderer) GetPreferredHeightForWidth(w WidgetLike, width int) (minimumHeight, naturalHeight int) {
	var m, n C.gint
	C.gtk_cell_renderer_get_preferred_height_for_width(self.object, w.W().object, C.gint(width), &m, &n)
	minimumHeight = int(m)
	naturalHeight = int(n)
	return
}

func (self *CellRenderer) GetPreferredSize(w WidgetLike) (minSize, natSize Requisition) {
	var m, n C.GtkRequisition
	C.gtk_cell_renderer_get_preferred_size(self.object, w.W().object, &m, &n)

	minSize.Height = int(m.height)
	minSize.Width = int(m.width)
	natSize.Height = int(n.height)
	natSize.Width = int(n.width)
	return
}

func (self *CellRenderer) GetPreferredWidth(w WidgetLike) (minSize, natSize int) {
	var m, n C.gint
	C.gtk_cell_renderer_get_preferred_width(self.object, w.W().object, &m, &n)
	minSize = int(m)
	natSize = int(n)
	return
}

func (self *CellRenderer) GetPreferredWidthForHeight(w WidgetLike, height int) (minWidth, natWidth int) {
	var m, n C.gint
	C.gtk_cell_renderer_get_preferred_width_for_height(self.object, w.W().object, C.gint(height), &m, &n)
	minWidth = int(m)
	natWidth = int(n)
	return
}

func (self *CellRenderer) GetRequestMode() int {
	return int(C.gtk_cell_renderer_get_request_mode(self.object))
}

//////////////////////////////
// END GtkCellRenderer
////////////////////////////// }}}

// GtkCellRendererText {{{
//////////////////////////////

// GtkCellRendererText type
type CellRendererText struct {
	object *C.GtkCellRendererText
	*CellRenderer
}

func NewCellRendererText() *CellRendererText {
	cl := &CellRendererText{}
	o := C.gtk_cell_renderer_text_new()
	cl.object = C.to_GtkCellRendererText(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(unsafe.Pointer(o)).(*CellRenderer)
	cellRendererTextFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererTextFinalizer(cl *CellRendererText) {
	runtime.SetFinalizer(cl, func(cl *CellRendererText) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererTextFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererText{}
	cl.object = C.to_GtkCellRendererText(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(obj).(*CellRenderer)
	cellRendererTextFinalizer(cl)

	return cl
}

func nativeFromCellRendererText(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererText)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_TEXT, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererText) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererText) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererText) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererText) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be CellRendererLike
func (self CellRendererText) CRenderer() *CellRenderer {
	return self.CellRenderer
}

// Interface functions

func (self *CellRendererText) SetFixedHeigthFromFont(numberOfRows int) {
	C.gtk_cell_renderer_text_set_fixed_height_from_font(self.object, C.gint(numberOfRows))
}

//////////////////////////////
// END GtkCellRendererText
////////////////////////////// }}}

// GtkCellRendererProgress {{{
//////////////////////////////

// GtkCellRendererProgress type
type CellRendererProgress struct {
	object *C.GtkCellRendererProgress
	*CellRenderer
	*Orientable
}

func NewCellRendererProgress() *CellRendererProgress {
	cl := &CellRendererProgress{}
	o := C.gtk_cell_renderer_progress_new()
	cl.object = C.to_GtkCellRendererProgress(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(unsafe.Pointer(o)).(*CellRenderer)
	cl.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	cellRendererProgressFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererProgressFinalizer(cl *CellRendererProgress) {
	runtime.SetFinalizer(cl, func(cl *CellRendererProgress) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererProgressFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererProgress{}
	cl.object = C.to_GtkCellRendererProgress(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(obj).(*CellRenderer)
	cl.Orientable = newOrientableFromNative(obj).(*Orientable)
	cellRendererProgressFinalizer(cl)

	return cl
}

func nativeFromCellRendererProgress(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererProgress)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_PROGRESS, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererProgress) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererProgress) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererProgress) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererProgress) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be CellRendererLike
func (self CellRendererProgress) CRenderer() *CellRenderer {
	return self.CellRenderer
}

// To be Orientable-like
func (self CellRendererProgress) IOrientable() *Orientable {
	return self.Orientable
}

//////////////////////////////
// END GtkCellRendererProgress
////////////////////////////// }}}

// GtkCellRendererSpinner {{{
//////////////////////////////

// GtkCellRendererSpinner type
type CellRendererSpinner struct {
	object *C.GtkCellRendererSpinner
	*CellRenderer
}

func NewCellRendererSpinner() *CellRendererSpinner {
	cl := &CellRendererSpinner{}
	o := C.gtk_cell_renderer_spinner_new()
	cl.object = C.to_GtkCellRendererSpinner(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(unsafe.Pointer(o)).(*CellRenderer)
	cellRendererSpinnerFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererSpinnerFinalizer(cl *CellRendererSpinner) {
	runtime.SetFinalizer(cl, func(cl *CellRendererSpinner) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererSpinnerFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererSpinner{}
	cl.object = C.to_GtkCellRendererSpinner(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(obj).(*CellRenderer)
	cellRendererSpinnerFinalizer(cl)

	return cl
}

func nativeFromCellRendererSpinner(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererSpinner)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_SPINNER, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererSpinner) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererSpinner) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererSpinner) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererSpinner) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be CellRendererLike
func (self CellRendererSpinner) CRenderer() *CellRenderer {
	return self.CellRenderer
}

//////////////////////////////
// END GtkCellRendererSpinner
////////////////////////////// }}}

// GtkCellRendererToggle {{{
//////////////////////////////

// GtkCellRendererToggle type
type CellRendererToggle struct {
	object *C.GtkCellRendererToggle
	*CellRenderer
}

func NewCellRendererToggle() *CellRendererToggle {
	cl := &CellRendererToggle{}
	o := C.gtk_cell_renderer_toggle_new()
	cl.object = C.to_GtkCellRendererToggle(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(unsafe.Pointer(o)).(*CellRenderer)
	cellRendererToggleFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererToggleFinalizer(cl *CellRendererToggle) {
	runtime.SetFinalizer(cl, func(cl *CellRendererToggle) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererToggleFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererToggle{}
	cl.object = C.to_GtkCellRendererToggle(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(obj).(*CellRenderer)
	cellRendererToggleFinalizer(cl)

	return cl
}

func nativeFromCellRendererToggle(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererToggle)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_TOGGLE, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererToggle) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererToggle) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererToggle) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererToggle) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be CellRendererLike
func (self CellRendererToggle) CRenderer() *CellRenderer {
	return self.CellRenderer
}

// CellRendererToggle Interface

func (self *CellRendererToggle) GetRadio() bool {
	b := C.gtk_cell_renderer_toggle_get_radio(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRendererToggle) SetRadio(radio bool) {
	b := gobject.GBool(radio)
	defer b.Free()
	C.gtk_cell_renderer_toggle_set_radio(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellRendererToggle) GetActive() bool {
	b := C.gtk_cell_renderer_toggle_get_active(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRendererToggle) SetActive(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_cell_renderer_toggle_set_active(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellRendererToggle) GetActivatable() bool {
	b := C.gtk_cell_renderer_toggle_get_activatable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRendererToggle) SetActivatable(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_cell_renderer_toggle_set_activatable(self.object, *((*C.gboolean)(b.GetPtr())))
}

//////////////////////////////
// END GtkCellRendererToggle
////////////////////////////// }}}

// GtkCellRendererPixbuf {{{
//////////////////////////////

// GtkCellRendererPixbuf type
type CellRendererPixbuf struct {
	object *C.GtkCellRendererPixbuf
	*CellRenderer
}

func NewCellRendererPixbuf() *CellRendererPixbuf {
	cl := &CellRendererPixbuf{}
	o := C.gtk_cell_renderer_pixbuf_new()
	cl.object = C.to_GtkCellRendererPixbuf(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(unsafe.Pointer(o)).(*CellRenderer)
	cellRendererPixbufFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererPixbufFinalizer(cl *CellRendererPixbuf) {
	runtime.SetFinalizer(cl, func(cl *CellRendererPixbuf) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererPixbufFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererPixbuf{}
	cl.object = C.to_GtkCellRendererPixbuf(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(obj).(*CellRenderer)
	cellRendererPixbufFinalizer(cl)

	return cl
}

func nativeFromCellRendererPixbuf(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererPixbuf)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_PIXBUF, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererPixbuf) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererPixbuf) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererPixbuf) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererPixbuf) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be CellRendererLike
func (self CellRendererPixbuf) CRenderer() *CellRenderer {
	return self.CellRenderer
}

//////////////////////////////
// END GtkCellRendererPixbuf
////////////////////////////// }}}

// GtkCellRendererAccel {{{
//////////////////////////////

// GtkCellRendererAccel type
type CellRendererAccel struct {
	object *C.GtkCellRendererAccel
	*CellRendererText
}

func NewCellRendererAccel() *CellRendererAccel {
	cl := &CellRendererAccel{}
	o := C.gtk_cell_renderer_accel_new()
	cl.object = C.to_GtkCellRendererAccel(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(unsafe.Pointer(o)).(*CellRendererText)
	cellRendererAccelFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererAccelFinalizer(cl *CellRendererAccel) {
	runtime.SetFinalizer(cl, func(cl *CellRendererAccel) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererAccelFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererAccel{}
	cl.object = C.to_GtkCellRendererAccel(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(obj).(*CellRendererText)
	cellRendererAccelFinalizer(cl)

	return cl
}

func nativeFromCellRendererAccel(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererAccel)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_ACCEL, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererAccel) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererAccel) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererAccel) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererAccel) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

//////////////////////////////
// END GtkCellRendererAccel
////////////////////////////// }}}

// GtkCellRendererCombo {{{
//////////////////////////////

// GtkCellRendererCombo type
type CellRendererCombo struct {
	object *C.GtkCellRendererCombo
	*CellRendererText
}

func NewCellRendererCombo() *CellRendererCombo {
	cl := &CellRendererCombo{}
	o := C.gtk_cell_renderer_combo_new()
	cl.object = C.to_GtkCellRendererCombo(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(unsafe.Pointer(o)).(*CellRendererText)
	cellRendererComboFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererComboFinalizer(cl *CellRendererCombo) {
	runtime.SetFinalizer(cl, func(cl *CellRendererCombo) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererComboFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererCombo{}
	cl.object = C.to_GtkCellRendererCombo(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(obj).(*CellRendererText)
	cellRendererComboFinalizer(cl)

	return cl
}

func nativeFromCellRendererCombo(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererCombo)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_COMBO, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererCombo) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererCombo) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererCombo) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererCombo) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

//////////////////////////////
// END GtkCellRendererCombo
////////////////////////////// }}}

// GtkCellRendererSpin {{{
//////////////////////////////

// GtkCellRendererSpin type
type CellRendererSpin struct {
	object *C.GtkCellRendererSpin
	*CellRendererText
}

func NewCellRendererSpin() *CellRendererSpin {
	cl := &CellRendererSpin{}
	o := C.gtk_cell_renderer_spin_new()
	cl.object = C.to_GtkCellRendererSpin(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(unsafe.Pointer(o)).(*CellRendererText)
	cellRendererSpinFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererSpinFinalizer(cl *CellRendererSpin) {
	runtime.SetFinalizer(cl, func(cl *CellRendererSpin) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererSpinFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererSpin{}
	cl.object = C.to_GtkCellRendererSpin(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(obj).(*CellRendererText)
	cellRendererSpinFinalizer(cl)

	return cl
}

func nativeFromCellRendererSpin(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererSpin)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_SPIN, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererSpin) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererSpin) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererSpin) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererSpin) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

//////////////////////////////
// END GtkCellRendererSpin
////////////////////////////// }}}

// GtkTreeViewColumn {{{
//////////////////////////////

// GtkTreeViewColumn type
type TreeViewColumn struct {
	object *C.GtkTreeViewColumn
}

func NewTreeViewColumn() *TreeViewColumn {
	tc := &TreeViewColumn{}
	tc.object = C.gtk_tree_view_column_new()

	if gobject.IsObjectFloating(tc) {
		gobject.RefSink(tc)
	}
	treeViewColumnFinalizer(tc)

	return tc
}

// TODO: gtk_tree_column_new_with_area

func NewTreeViewColumnWithAttributes(title string, cell CellRendererLike, attr A) *TreeViewColumn {
	tc := NewTreeViewColumn()
	tc.SetTitle(title)
	tc.PackStart(cell, true)
	tc.SetAttributes(cell, attr)

	return tc
}

// Clear TreeViewColumn struct when it goes out of reach
func treeViewColumnFinalizer(tc *TreeViewColumn) {
	runtime.SetFinalizer(tc, func(tc *TreeViewColumn) { gobject.Unref(tc) })
}

// Conversion functions
func newTreeViewColumnFromNative(obj unsafe.Pointer) interface{} {
	tc := &TreeViewColumn{}
	tc.object = C.to_GtkTreeViewColumn(obj)

	if gobject.IsObjectFloating(tc) {
		gobject.RefSink(tc)
	} else {
		gobject.Ref(tc)
	}
	treeViewColumnFinalizer(tc)

	return tc
}

func nativeFromTreeViewColumn(tc interface{}) *gobject.GValue {
	tvc, ok := tc.(*TreeViewColumn)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TREE_VIEW_COLUMN, tvc.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TreeViewColumn) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TreeViewColumn) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TreeViewColumn) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TreeViewColumn) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// TreeViewColumn interface

func (self *TreeViewColumn) PackStart(cell CellRendererLike, expand bool) {
	b := gobject.GBool(expand)
	defer b.Free()
	C.gtk_tree_view_column_pack_start(self.object, cell.CRenderer().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) PackEnd(cell CellRendererLike, expand bool) {
	b := gobject.GBool(expand)
	defer b.Free()
	C.gtk_tree_view_column_pack_start(self.object, cell.CRenderer().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) Clear() {
	C.gtk_tree_view_column_clear(self.object)
}

func (self *TreeViewColumn) AddAttribute(cellRenderer CellRendererLike, attribute string, column int) {
	s := gobject.GString(attribute)
	defer s.Free()
	C.gtk_tree_view_column_add_attribute(self.object, cellRenderer.CRenderer().object, (*C.gchar)(s.GetPtr()), C.gint(column))
}

func (self *TreeViewColumn) SetAttributes(cellRenderer CellRendererLike, attrs A) {
	self.ClearAttributes(cellRenderer)
	for _, cellAttr := range attrs {
		self.AddAttribute(cellRenderer, cellAttr.Attribute, cellAttr.Column)
	}
}

//TODO: gtk_tree_view_column_set_data_func

func (self *TreeViewColumn) ClearAttributes(cellRenderer CellRendererLike) {
	C.gtk_tree_view_column_clear_attributes(self.object, cellRenderer.CRenderer().object)
}

func (self *TreeViewColumn) SetSpacing(spacing int) {
	C.gtk_tree_view_column_set_spacing(self.object, C.gint(spacing))
}

func (self *TreeViewColumn) GetSpacing() int {
	return int(C.gtk_tree_view_column_get_spacing(self.object))
}

func (self *TreeViewColumn) SetVisible(visible bool) {
	b := gobject.GBool(visible)
	defer b.Free()
	C.gtk_tree_view_column_set_visible(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetVisible() bool {
	b := C.gtk_tree_view_column_get_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetResizable(resizable bool) {
	b := gobject.GBool(resizable)
	defer b.Free()
	C.gtk_tree_view_column_set_resizable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetResizable() bool {
	b := C.gtk_tree_view_column_get_resizable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetSizing(gtk_TreeViewColumnSizing int) {
	C.gtk_tree_view_column_set_sizing(self.object, C.GtkTreeViewColumnSizing(gtk_TreeViewColumnSizing))
}

func (self *TreeViewColumn) GetSizing() int {
	return int(C.gtk_tree_view_column_get_sizing(self.object))
}

func (self *TreeViewColumn) GetWidth() int {
	return int(C.gtk_tree_view_column_get_width(self.object))
}

func (self *TreeViewColumn) GetFixedWidth() int {
	return int(C.gtk_tree_view_column_get_fixed_width(self.object))
}

func (self *TreeViewColumn) SetFixedWidth(fixedWidth int) {
	C.gtk_tree_view_column_set_fixed_width(self.object, C.gint(fixedWidth))
}

func (self *TreeViewColumn) SetMinWidth(minWidth int) {
	C.gtk_tree_view_column_set_min_width(self.object, C.gint(minWidth))
}

func (self *TreeViewColumn) GetMinWidth() int {
	return int(C.gtk_tree_view_column_get_min_width(self.object))
}

func (self *TreeViewColumn) SetMaxWidth(maxWidth int) {
	C.gtk_tree_view_column_set_max_width(self.object, C.gint(maxWidth))
}

func (self *TreeViewColumn) GetMaxWidth() int {
	return int(C.gtk_tree_view_column_get_max_width(self.object))
}

func (self *TreeViewColumn) Clicked() {
	C.gtk_tree_view_column_clicked(self.object)
}

func (self *TreeViewColumn) SetTitle(title string) {
	ctitle := gobject.GString(title)
	defer ctitle.Free()
	C.gtk_tree_view_column_set_title(self.object, (*C.gchar)(ctitle.GetPtr()))
}

func (self *TreeViewColumn) GetTitle() string {
	s := C.gtk_tree_view_column_get_title(self.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *TreeViewColumn) SetExpand(expand bool) {
	b := gobject.GBool(expand)
	defer b.Free()
	C.gtk_tree_view_column_set_expand(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetExpand() bool {
	b := C.gtk_tree_view_column_get_expand(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetClickable(clickable bool) {
	b := gobject.GBool(clickable)
	defer b.Free()
	C.gtk_tree_view_column_set_clickable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetClickable() bool {
	b := C.gtk_tree_view_column_get_clickable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetWidget(w WidgetLike) {
	var cw *C.GtkWidget = nil
	if w != nil {
		cw = w.W().object
	}
	C.gtk_tree_view_column_set_widget(self.object, cw)
}

func (self *TreeViewColumn) GetWidget() WidgetLike {
	w := C.gtk_tree_view_column_get_widget(self.object)
	if w == nil {
		return nil
	}

	wid, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

func (self *TreeViewColumn) GetButton() WidgetLike {
	w := C.gtk_tree_view_column_get_button(self.object)

	wid, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

func (self *TreeViewColumn) SetAlignment(xalign float32) {
	C.gtk_tree_view_column_set_alignment(self.object, C.gfloat(xalign))
}

func (self *TreeViewColumn) GetAlignment() float32 {
	return float32(C.gtk_tree_view_column_get_alignment(self.object))
}

func (self *TreeViewColumn) SetReorderable(reorderable bool) {
	b := gobject.GBool(reorderable)
	defer b.Free()
	C.gtk_tree_view_column_set_reorderable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetReorderable() bool {
	b := C.gtk_tree_view_column_get_reorderable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetSortColumnId(sortColumnId int) {
	C.gtk_tree_view_column_set_sort_column_id(self.object, C.gint(sortColumnId))
}

func (self *TreeViewColumn) GetSortColumnId() int {
	return int(C.gtk_tree_view_column_get_sort_column_id(self.object))
}

func (self *TreeViewColumn) SetSortIndicator(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_tree_view_column_set_sort_indicator(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetSortIndicator() bool {
	b := C.gtk_tree_view_column_get_sort_indicator(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetSortOrder(gtk_SortType int) {
	C.gtk_tree_view_column_set_sort_order(self.object, C.GtkSortType(gtk_SortType))
}

func (self *TreeViewColumn) GetSortOrder() int {
	return int(C.gtk_tree_view_column_get_sort_order(self.object))
}

func (self *TreeViewColumn) CellSetCellData(model TreeModelLike, iter *TreeIter, isExpander, isExpanded bool) {
	b1 := gobject.GBool(isExpander)
	defer b1.Free()
	b2 := gobject.GBool(isExpanded)
	defer b2.Free()
	C.gtk_tree_view_column_cell_set_cell_data(self.object, model.ITreeModel().object, &iter.object, *((*C.gboolean)(b1.GetPtr())),
		*((*C.gboolean)(b2.GetPtr())))
}

func (self *TreeViewColumn) CellGetSize(rec gdk3.Rectangle) (xOffset, yOffset, width, height int) {
	rectangle := gobject.ConvertToC(rec)
	defer rectangle.Free()

	var x, y, w, h C.gint
	C.gtk_tree_view_column_cell_get_size(self.object, (*C.GdkRectangle)(rectangle.GetPtr()), &x, &y, &w, &h)
	xOffset = int(x)
	yOffset = int(y)
	width = int(w)
	height = int(h)
	return
}

func (self *TreeViewColumn) CellGetPosition(renderer CellRendererLike) (bool, int, int) {
	var o, w C.gint
	b := C.gtk_tree_view_column_cell_get_position(self.object, renderer.CRenderer().object, &o, &w)
	retBool := gobject.GoBool(unsafe.Pointer(&b))
	return retBool, int(o), int(w)
}

func (self *TreeViewColumn) CellIsVisible() bool {
	b := C.gtk_tree_view_column_cell_is_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) FocusCell(renderer CellRendererLike) {
	C.gtk_tree_view_column_focus_cell(self.object, renderer.CRenderer().object)
}

func (self *TreeViewColumn) QueueResize() {
	C.gtk_tree_view_column_queue_resize(self.object)
}

func (self *TreeViewColumn) GetTreeView() WidgetLike {
	w := C.gtk_tree_view_column_get_tree_view(self.object)
	if w == nil {
		return nil
	}

	widget, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return widget.(WidgetLike)
	}
	return nil
}

func (self *TreeViewColumn) GetXOffset() int {
	return int(C.gtk_tree_view_column_get_x_offset(self.object))
}

//////////////////////////////
// End GtkTreeViewColumn
////////////////////////////// }}}

// GtkTreeSelection {{{
//////////////////////////////

// GtkTreeSelection type
type TreeSelection struct {
	object *C.GtkTreeSelection
}

// Clear TreeSelection struct when it goes out of reach
func treeSelectionFinalizer(ts *TreeSelection) {
	runtime.SetFinalizer(ts, func(ts *TreeSelection) { gobject.Unref(ts) })
}

// Conversion function
func newTreeSelectionFromNative(obj unsafe.Pointer) interface{} {
	ts := &TreeSelection{}
	ts.object = C.to_GtkTreeSelection(obj)

	if gobject.IsObjectFloating(ts) {
		gobject.RefSink(ts)
	} else {
		gobject.Ref(ts)
	}
	treeSelectionFinalizer(ts)

	return ts
}

func nativeFromTreeSelection(ts interface{}) *gobject.GValue {
	treeSelection, ok := ts.(*TreeSelection)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TREE_SELECTION, treeSelection.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TreeSelection) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TreeSelection) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TreeSelection) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TreeSelection) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// TreeSelection interface

func (self *TreeSelection) SetMode(gtk_SelectionMode int) {
	C.gtk_tree_selection_set_mode(self.object, C.GtkSelectionMode(gtk_SelectionMode))
}

func (self *TreeSelection) GetMode() int {
	return int(C.gtk_tree_selection_get_mode(self.object))
}

func (self *TreeSelection) SetSelectFunction(f interface{}, data ...interface{}) {
	if f == nil {
		C._gtk_tree_select_set_select_function(self.object, 0)
		return
	}

	call, id := gobject.CreateCustomClosure(f, data...)
	_closures[id] = call
	C._gtk_tree_select_set_select_function(self.object, C.gint64(id))
}

func (self *TreeSelection) GetTreeView() *TreeView {
	w := C.gtk_tree_selection_get_tree_view(self.object)
	if w != nil {
		if view, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
			return view.(*TreeView)
		}
	}
	return nil
}

func (self *TreeSelection) GetSelected(iter *TreeIter) {
	var cIter *C.GtkTreeIter
	if iter != nil {
		cIter = &iter.object
	}
	C.gtk_tree_selection_get_selected(self.object, nil, cIter)
}

//TODO: gtk_tree_selection_selected_foreach

func (self *TreeSelection) GetSelectedRows() *glib.GList {
	list := C.gtk_tree_selection_get_selected_rows(self.object, nil)

	goList := glib.NewGListFromNative(unsafe.Pointer(list))
	goList.GC_FreeFull = true
	goList.DestroyFunc = func(obj unsafe.Pointer) { path := (*C.GtkTreePath)(obj); C.gtk_tree_path_free(path) }
	goList.ConversionFunc = newTreePathFromNative
	glib.GListFinalizer(goList)

	return goList
}

func (self *TreeSelection) CountSelectedRows() int {
	return int(C.gtk_tree_selection_count_selected_rows(self.object))
}

func (self *TreeSelection) SelectPath(path *TreePath) {
	C.gtk_tree_selection_select_path(self.object, path.object)
}

func (self *TreeSelection) UnselectSelectPath(path *TreePath) {
	C.gtk_tree_selection_unselect_path(self.object, path.object)
}

func (self *TreeSelection) PathIsSelected(path *TreePath) bool {
	b := C.gtk_tree_selection_path_is_selected(self.object, path.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeSelection) SelectIter(iter *TreeIter) {
	C.gtk_tree_selection_select_iter(self.object, &iter.object)
}

func (self *TreeSelection) UnselectIter(iter *TreeIter) {
	C.gtk_tree_selection_unselect_iter(self.object, &iter.object)
}

func (self *TreeSelection) IterIsSelected(iter *TreeIter) bool {
	b := C.gtk_tree_selection_iter_is_selected(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeSelection) SelectAll() {
	C.gtk_tree_selection_select_all(self.object)
}

func (self *TreeSelection) UnselectAll() {
	C.gtk_tree_selection_unselect_all(self.object)
}

func (self *TreeSelection) SelectRange(startPath, endPath *TreePath) {
	C.gtk_tree_selection_select_range(self.object, startPath.object, endPath.object)
}

func (self *TreeSelection) UnselectRange(startPath, endPath *TreePath) {
	C.gtk_tree_selection_unselect_range(self.object, startPath.object, endPath.object)
}

//////////////////////////////
// End GtkTreeSelection
////////////////////////////// }}}

// GtkTreeView {{{
//////////////////////////////

// GtkTreeView type
type TreeView struct {
	object *C.GtkTreeView
	*Container
}

func NewTreeView() *TreeView {
	tw := &TreeView{}
	o := C.gtk_tree_view_new()
	tw.object = C.to_GtkTreeView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(tw) {
		gobject.RefSink(tw)
	}
	tw.Container = NewContainer(unsafe.Pointer(o))
	treeViewFinalizer(tw)

	return tw
}

func NewTreeViewWithModel(model TreeModelLike) *TreeView {
	tw := &TreeView{}
	o := C.gtk_tree_view_new_with_model(model.ITreeModel().object)
	tw.object = C.to_GtkTreeView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(tw) {
		gobject.RefSink(tw)
	}
	tw.Container = NewContainer(unsafe.Pointer(o))
	treeViewFinalizer(tw)

	return tw
}

// Clear TreeView struct when it goes out of reach
func treeViewFinalizer(tw *TreeView) {
	runtime.SetFinalizer(tw, func(tw *TreeView) { gobject.Unref(tw) })
}

// Conversion functions
func newTreeViewFromNative(obj unsafe.Pointer) interface{} {
	tw := &TreeView{}
	tw.object = C.to_GtkTreeView(obj)

	if gobject.IsObjectFloating(tw) {
		gobject.RefSink(tw)
	} else {
		gobject.Ref(tw)
	}
	tw.Container = NewContainer(obj)
	treeViewFinalizer(tw)

	return tw
}

func nativeFromTreeView(tw interface{}) *gobject.GValue {
	treeV, ok := tw.(*TreeView)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TREE_VIEW, treeV.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TreeView) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TreeView) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TreeView) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TreeView) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self TreeView) C() *Container {
	return self.Container
}

// TreeView interface

func (self *TreeView) GetLevelIndentation() int {
	return int(C.gtk_tree_view_get_level_indentation(self.object))
}

func (self *TreeView) GetShowExpanders() bool {
	b := C.gtk_tree_view_get_show_expanders(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetLevelIndentation(indentation int) {
	C.gtk_tree_view_set_level_indentation(self.object, C.gint(indentation))
}

func (self *TreeView) SetShowExpanders(enabled bool) {
	b := gobject.GBool(enabled)
	defer b.Free()
	C.gtk_tree_view_set_show_expanders(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) GetModel() TreeModelLike {
	m := C.gtk_tree_view_get_model(self.object)
	if m == nil {
		return nil
	}
	model, err := gobject.ConvertToGo(unsafe.Pointer(m))

	if err == nil {
		return model.(TreeModelLike)
	}
	return nil
}

func (self *TreeView) SetModel(model TreeModelLike) {
	C.gtk_tree_view_set_model(self.object, model.ITreeModel().object)
}

func (self *TreeView) GetSelection() *TreeSelection {
	s := C.gtk_tree_view_get_selection(self.object)

	sel, err := gobject.ConvertToGo(unsafe.Pointer(s))
	if err == nil {
		return sel.(*TreeSelection)
	}
	return nil
}

func (self *TreeView) GetHeadersVisible() bool {
	b := C.gtk_tree_view_get_headers_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetHeadersVisible(headersVisible bool) {
	b := gobject.GBool(headersVisible)
	defer b.Free()
	C.gtk_tree_view_set_headers_visible(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) ColumnsAutoSize() {
	C.gtk_tree_view_columns_autosize(self.object)
}

func (self *TreeView) GetHeadersClickable() bool {
	b := C.gtk_tree_view_get_headers_clickable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetHeadersClickable(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_tree_view_set_headers_clickable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) SetRulesHint(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_tree_view_set_rules_hint(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) GetRulesHint() bool {
	b := C.gtk_tree_view_get_rules_hint(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) AppendColumn(column *TreeViewColumn) int {
	return int(C.gtk_tree_view_append_column(self.object, column.object))
}

func (self *TreeView) RemoveColumn(colum *TreeViewColumn) int {
	return int(C.gtk_tree_view_remove_column(self.object, colum.object))
}

func (self *TreeView) InsertColumn(colum *TreeViewColumn, position int) int {
	return int(C.gtk_tree_view_insert_column(self.object, colum.object, C.gint(position)))
}

func (self *TreeView) InsertColumnWithAttributes(position int, title string, cell CellRendererLike, attrs A) int {
	col := NewTreeViewColumnWithAttributes(title, cell, attrs)
	return int(self.InsertColumn(col, position))
}

//TODO: gtk_tree_view_insert_column_with_data_func

func (self *TreeView) GetColumn(n int) *TreeViewColumn {
	tc := C.gtk_tree_view_get_column(self.object, C.gint(n))
	if tc == nil {
		return nil
	}

	column, err := gobject.ConvertToGo(unsafe.Pointer(tc))
	if err == nil {
		return column.(*TreeViewColumn)
	}
	return nil
}

func (self *TreeView) GetColumns() *glib.GList {
	l := C.gtk_tree_view_get_columns(self.object)

	if l != nil {
		goList := glib.NewGListFromNative(unsafe.Pointer(l))
		goList.GC_Free = true
		goList.GC_FreeFull = false
		goList.ConversionFunc = newTreeViewColumnFromNative
		glib.GListFinalizer(goList)
		return goList
	}
	return nil
}

func (self *TreeView) MoveColumnAfter(column, baseColumn *TreeViewColumn) {
	var base *C.GtkTreeViewColumn = nil
	if baseColumn != nil {
		base = baseColumn.object
	}
	C.gtk_tree_view_move_column_after(self.object, column.object, base)
}

func (self *TreeView) SetExpanderColumn(column *TreeViewColumn) {
	var col *C.GtkTreeViewColumn = nil
	if column != nil {
		col = column.object
	}
	C.gtk_tree_view_set_expander_column(self.object, col)
}

func (self *TreeView) GetExpanderColumn() *TreeViewColumn {
	ec := C.gtk_tree_view_get_expander_column(self.object)

	if ec != nil {
		expanderCol, err := gobject.ConvertToGo(unsafe.Pointer(ec))
		if err == nil {
			return expanderCol.(*TreeViewColumn)
		}
	}
	return nil
}

//TODO: gtk_tree_view_set_column_drag_function

func (self *TreeView) ScrollToPoint(xTree, yTree int) {
	C.gtk_tree_view_scroll_to_point(self.object, C.gint(xTree), C.gint(yTree))
}

func (self *TreeView) ScrollToCell(path *TreePath, column *TreeViewColumn, useAlign bool, rowAlign, colAlign float32) {
	var p *C.GtkTreePath = nil
	var c *C.GtkTreeViewColumn = nil
	b := gobject.GBool(useAlign)
	defer b.Free()

	if path != nil {
		p = path.object
	}

	if column != nil {
		c = column.object
	}
	C.gtk_tree_view_scroll_to_cell(self.object, p, c, *((*C.gboolean)(b.GetPtr())), C.gfloat(rowAlign), C.gfloat(colAlign))
}

func (self *TreeView) SetCursor(path *TreePath, focusColumn *TreeViewColumn, startEditing bool) {
	var c *C.GtkTreeViewColumn = nil
	if focusColumn != nil {
		c = focusColumn.object
	}

	b := gobject.GBool(startEditing)
	defer b.Free()

	C.gtk_tree_view_set_cursor(self.object, path.object, c, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) SetCursorOnCell(path *TreePath, focusColumn *TreeViewColumn, focusCell CellRendererLike, startEditing bool) {
	var c *C.GtkTreeViewColumn = nil
	var r *C.GtkCellRenderer = nil

	if focusColumn != nil {
		c = focusColumn.object
	}

	if focusCell != nil {
		r = focusCell.CRenderer().object
	}

	b := gobject.GBool(startEditing)
	defer b.Free()
	C.gtk_tree_view_set_cursor_on_cell(self.object, path.object, c, r, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) GetCursor() (path *TreePath, focusCol *TreeViewColumn) {
	path, focusCol = nil, nil
	var cpath *C.GtkTreePath
	var ccol *C.GtkTreeViewColumn
	C.gtk_tree_view_get_cursor(self.object, &cpath, &ccol)

	if cpath != nil {
		path = newTreePathFromNative(unsafe.Pointer(cpath)).(*TreePath)
	}

	if ccol != nil {
		focusCol = newTreeViewColumnFromNative(unsafe.Pointer(ccol)).(*TreeViewColumn)
	}
	return
}

func (self *TreeView) RowActivated(path *TreePath, column *TreeViewColumn) {
	C.gtk_tree_view_row_activated(self.object, path.object, column.object)
}

func (self *TreeView) ExpandAll() {
	C.gtk_tree_view_expand_all(self.object)
}

func (self *TreeView) CollapseAll() {
	C.gtk_tree_view_collapse_all(self.object)
}

func (self *TreeView) ExpandRow(path *TreePath, openAll bool) {
	b := gobject.GBool(openAll)
	defer b.GetPtr()
	C.gtk_tree_view_expand_row(self.object, path.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) CollapseRow(path *TreePath) bool {
	b := C.gtk_tree_view_collapse_row(self.object, path.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//TODO: gtk_tree_view_map_expanded_rows

func (self *TreeView) RowExpanded(path *TreePath) bool {
	b := C.gtk_tree_view_row_expanded(self.object, path.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetReorderable(reorderable bool) {
	b := gobject.GBool(reorderable)
	defer b.Free()
	C.gtk_tree_view_set_reorderable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) GetReorderable() bool {
	b := C.gtk_tree_view_get_reorderable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//TODO: gtk_tree_view_get_path_at_pos
//TODO: gtk_tree_view_is_blank_at_pos

func (self *TreeView) GetCellArea(path *TreePath, column *TreeViewColumn) gdk3.Rectangle {
	var cpath *C.GtkTreePath = nil
	var ccolumn *C.GtkTreeViewColumn = nil
	var rec C.GdkRectangle

	if path != nil {
		cpath = path.object
	}

	if column != nil {
		ccolumn = column.object
	}
	C.gtk_tree_view_get_cell_area(self.object, cpath, ccolumn, &rec)

	if goRec, err := gobject.ConvertToGo(unsafe.Pointer(&rec), gdk3.GdkType.RECTANGLE); err == nil {
		return goRec.(gdk3.Rectangle)
	}
	return gdk3.Rectangle{}
}

func (self *TreeView) GetBackgroundArea(path *TreePath, column *TreeViewColumn) gdk3.Rectangle {
	var cpath *C.GtkTreePath = nil
	var ccolumn *C.GtkTreeViewColumn = nil
	var rec C.GdkRectangle

	if path != nil {
		cpath = path.object
	}

	if column != nil {
		ccolumn = column.object
	}
	C.gtk_tree_view_get_background_area(self.object, cpath, ccolumn, &rec)

	if goRec, err := gobject.ConvertToGo(unsafe.Pointer(&rec), gdk3.GdkType.RECTANGLE); err == nil {
		return goRec.(gdk3.Rectangle)
	}
	return gdk3.Rectangle{}
}

func (self *TreeView) GetVisibleRect() gdk3.Rectangle {
	var rec C.GdkRectangle

	C.gtk_tree_view_get_visible_rect(self.object, &rec)

	if goRec, err := gobject.ConvertToGo(unsafe.Pointer(&rec), gdk3.GdkType.RECTANGLE); err == nil {
		return goRec.(gdk3.Rectangle)
	}
	return gdk3.Rectangle{}
}

func (self *TreeView) GetVisibleRange() (pathsValid bool, start, end *TreePath) {
	start, end = nil, nil
	var cstart, cend *C.GtkTreePath
	b := C.gtk_tree_view_get_visible_range(self.object, &cstart, &cend)
	pathsValid = gobject.GoBool(unsafe.Pointer(&b))

	if cstart != nil {
		start = newTreePathFromNative(unsafe.Pointer(cstart)).(*TreePath)
	}

	if cend != nil {
		end = newTreePathFromNative(unsafe.Pointer(cend)).(*TreePath)
	}
	return
}

func (self *TreeView) GetBinWindow() *gdk3.Window {
	w := C.gtk_tree_view_get_bin_window(self.object)
	if w != nil {
		if win, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
			return win.(*gdk3.Window)
		}
	}
	return nil
}

//TODO: gtk_tree_view_convert_bin_window_to_tree_coords
//TODO: gtk_tree_view_convert_bin_window_to_widget_coords
//TODO: gtk_tree_view_convert_tree_to_widget_coords
//TODO: gtk_tree_view_convert_tree_to_bin_window_coords

func (self *TreeView) SetEnableSearch(enableSearch bool) {
	b := gobject.GBool(enableSearch)
	defer b.Free()
	C.gtk_tree_view_set_enable_search(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) GetEnableSearch() bool {
	b := C.gtk_tree_view_get_enable_search(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) GetSearchColumn() int {
	return int(C.gtk_tree_view_get_search_column(self.object))
}

func (self *TreeView) SetSearchColumn(column int) {
	C.gtk_tree_view_set_search_column(self.object, C.gint(column))
}

//TODO: gtk_tree_view_set_search_equal_func

func (self *TreeView) GetSearchEntry() *Entry {
	e := C.gtk_tree_view_get_search_entry(self.object)
	if e != nil {
		return newEntryFromNative(unsafe.Pointer(e)).(*Entry)
	}
	return nil
}

func (self *TreeView) SetSearchEntry(entry *Entry) {
	if entry == nil {
		C.gtk_tree_view_set_search_entry(self.object, nil)
	}
	C.gtk_tree_view_set_search_entry(self.object, entry.object)
}

//TODO: gtk_tree_view_set_search_position_func

func (self *TreeView) GetFixedHeightMode() bool {
	b := C.gtk_tree_view_get_fixed_height_mode(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetFixedHeightMode(enable bool) {
	b := gobject.GBool(enable)
	defer b.Free()
	C.gtk_tree_view_set_fixed_height_mode(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) GetHoverSelection() bool {
	b := C.gtk_tree_view_get_hover_selection(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetHoverSelection(hover bool) {
	b := gobject.GBool(hover)
	defer b.Free()
	C.gtk_tree_view_set_hover_selection(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) GetHoverExpand() bool {
	b := C.gtk_tree_view_get_hover_expand(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetHoverExpand(expand bool) {
	b := gobject.GBool(expand)
	defer b.Free()
	C.gtk_tree_view_set_hover_expand(self.object, *((*C.gboolean)(b.GetPtr())))
}

//TODO: gtk_tree_view_set_row_seperator_func

func (self *TreeView) GetRubberBanding() bool {
	b := C.gtk_tree_view_get_rubber_banding(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetRubberBanding(expand bool) {
	b := gobject.GBool(expand)
	defer b.Free()
	C.gtk_tree_view_set_rubber_banding(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) IsRubberBandingActive() bool {
	b := C.gtk_tree_view_is_rubber_banding_active(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) GetEnableTreeLines() bool {
	b := C.gtk_tree_view_get_enable_tree_lines(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetEnableTreeLines(enabled bool) {
	b := gobject.GBool(enabled)
	defer b.Free()
	C.gtk_tree_view_set_enable_tree_lines(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) GetGridLines() int {
	return int(C.gtk_tree_view_get_grid_lines(self.object))
}

func (self *TreeView) SetGridLines(gtk_TreeViewGridLines int) {
	C.gtk_tree_view_set_grid_lines(self.object, C.GtkTreeViewGridLines(gtk_TreeViewGridLines))
}

//TODO: set_tooltip_row
//TODO: set_tooltip_cell
//TODO: get_tooltip_context

func (self *TreeView) GetTooltipColumn() int {
	return int(C.gtk_tree_view_get_tooltip_column(self.object))
}

func (self *TreeView) SetTooltipColumn(column int) {
	C.gtk_tree_view_set_tooltip_column(self.object, C.gint(column))
}

//////////////////////////////
// End GtkTreeView
////////////////////////////// }}}

// End Tree, List and Icon Grid Widgets }}}

// Menus, Combo Box, Toolbar {{{

// GtkComboBox {{{
//////////////////////////////

// GtkComboBox type
type ComboBox struct {
	object *C.GtkComboBox
	*Bin
	*CellLayout
}

func NewComboBox() *ComboBox {
	cb := &ComboBox{}
	o := C.gtk_combo_box_new()
	cb.object = C.to_GtkComboBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.Bin = NewBin(unsafe.Pointer(o))
	cb.CellLayout = newCellLayoutFromNative(unsafe.Pointer(o)).(*CellLayout)
	comboBoxFinalizer(cb)

	return cb
}

func NewComboBoxWithEntry() *ComboBox {
	cb := &ComboBox{}
	o := C.gtk_combo_box_new_with_entry()
	cb.object = C.to_GtkComboBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.Bin = NewBin(unsafe.Pointer(o))
	cb.CellLayout = newCellLayoutFromNative(unsafe.Pointer(o)).(*CellLayout)
	comboBoxFinalizer(cb)

	return cb
}

func NewComboBoxWithModel(model TreeModelLike) *ComboBox {
	cb := &ComboBox{}
	o := C.gtk_combo_box_new_with_model(model.ITreeModel().object)
	cb.object = C.to_GtkComboBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.Bin = NewBin(unsafe.Pointer(o))
	cb.CellLayout = newCellLayoutFromNative(unsafe.Pointer(o)).(*CellLayout)
	comboBoxFinalizer(cb)

	return cb
}

func NewComboBoxWithModelAndEntry(model TreeModelLike) *ComboBox {
	cb := &ComboBox{}
	o := C.gtk_combo_box_new_with_model_and_entry(model.ITreeModel().object)
	cb.object = C.to_GtkComboBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.Bin = NewBin(unsafe.Pointer(o))
	cb.CellLayout = newCellLayoutFromNative(unsafe.Pointer(o)).(*CellLayout)
	comboBoxFinalizer(cb)

	return cb
}

func NewComboBoxWithArea(area CellAreaLike) *ComboBox {
	cb := &ComboBox{}
	o := C.gtk_combo_box_new_with_area(area.CArea().object)
	cb.object = C.to_GtkComboBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.Bin = NewBin(unsafe.Pointer(o))
	cb.CellLayout = newCellLayoutFromNative(unsafe.Pointer(o)).(*CellLayout)
	comboBoxFinalizer(cb)

	return cb
}

func NewComboBoxWithAreaAndEntry(area CellAreaLike) *ComboBox {
	cb := &ComboBox{}
	o := C.gtk_combo_box_new_with_area_and_entry(area.CArea().object)
	cb.object = C.to_GtkComboBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.Bin = NewBin(unsafe.Pointer(o))
	cb.CellLayout = newCellLayoutFromNative(unsafe.Pointer(o)).(*CellLayout)
	comboBoxFinalizer(cb)

	return cb
}

// Clear ComboBox struct when it goes out of reach
func comboBoxFinalizer(cb *ComboBox) {
	runtime.SetFinalizer(cb, func(cb *ComboBox) { gobject.Unref(cb) })
}

// Conversion funcs
func newComboBoxFromNative(obj unsafe.Pointer) interface{} {
	cb := &ComboBox{}
	cb.object = C.to_GtkComboBox(obj)

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	} else {
		gobject.Ref(cb)
	}
	cb.Bin = NewBin(obj)
	cb.CellLayout = newCellLayoutFromNative(obj).(*CellLayout)
	comboBoxFinalizer(cb)

	return cb
}

func nativeFromComboBox(cb interface{}) *gobject.GValue {
	combo, ok := cb.(*ComboBox)
	if ok {
		gv := gobject.CreateCGValue(GtkType.COMBO_BOX, combo.ToNative())
		return gv
	}
	return nil
}

// To be Object-like
func (self ComboBox) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ComboBox) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ComboBox) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ComboBox) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be bin-like
func (self ComboBox) CBin() *Bin {
	return self.Bin
}

// Implement CellLayoutLike
func (self ComboBox) ICellLayout() *CellLayout {
	return self.CellLayout
}

// ComboBox interface

func (self *ComboBox) GetWrapWidth() int {
	return int(C.gtk_combo_box_get_wrap_width(self.object))
}

func (self *ComboBox) SetWrapWidth(width int) {
	C.gtk_combo_box_set_wrap_width(self.object, C.gint(width))
}

func (self *ComboBox) GetRowSpanColumn() int {
	return int(C.gtk_combo_box_get_row_span_column(self.object))
}

func (self *ComboBox) SetRowSpanColumn(rowSpan int) {
	C.gtk_combo_box_set_row_span_column(self.object, C.gint(rowSpan))
}

func (self *ComboBox) GetColumnSpanColumn() int {
	return int(C.gtk_combo_box_get_column_span_column(self.object))
}

func (self *ComboBox) SetColumnSpanColumn(columnSpan int) {
	C.gtk_combo_box_set_column_span_column(self.object, C.gint(columnSpan))
}

func (self *ComboBox) GetActive() int {
	return int(C.gtk_combo_box_get_active(self.object))
}

func (self *ComboBox) SetActive(index int) {
	C.gtk_combo_box_set_active(self.object, C.gint(index))
}

func (self *ComboBox) GetActiveIter(iter *TreeIter) bool {
	b := C.gtk_combo_box_get_active_iter(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ComboBox) SetActiveIter(iter *TreeIter) {
	C.gtk_combo_box_set_active_iter(self.object, &iter.object)
}

func (self *ComboBox) GetIdColumn() int {
	return int(C.gtk_combo_box_get_id_column(self.object))
}

func (self *ComboBox) SetIdColumn(idColumn int) {
	C.gtk_combo_box_set_id_column(self.object, C.gint(idColumn))
}

func (self *ComboBox) GetActiveId() string {
	s := C.gtk_combo_box_get_active_id(self.object)
	if s == nil {
		return ""
	}
	sd := C.g_strdup(s)
	return gobject.GoString(unsafe.Pointer(sd))
}

func (self *ComboBox) SetActiveId(activeId interface{}) bool {
	if activeId == nil {
		C.gtk_combo_box_set_active_id(self.object, nil)
		return true
	}

	if s, ok := activeId.(string); ok {
		st := gobject.GString(s)
		defer st.Free()
		b := C.gtk_combo_box_set_active_id(self.object, (*C.gchar)(st.GetPtr()))
		return gobject.GoBool(unsafe.Pointer(&b))
	}
	return false
}

func (self *ComboBox) GetModel() TreeModelLike {
	m := C.gtk_combo_box_get_model(self.object)
	if m == nil {
		return nil
	}

	if model, err := gobject.ConvertToGo(unsafe.Pointer(m)); err == nil {
		return model.(TreeModelLike)
	}
	return nil
}

func (self *ComboBox) SetModel(model TreeModelLike) {
	if model == nil {
		C.gtk_combo_box_set_model(self.object, nil)
	}
	C.gtk_combo_box_set_model(self.object, model.ITreeModel().object)
}

func (self *ComboBox) PopupForDevice(device *gdk3.Device) {
	C.gtk_combo_box_popup_for_device(self.object, (*C.GdkDevice)(device.ToNative()))
}

func (self *ComboBox) Popup() {
	C.gtk_combo_box_popup(self.object)
}

func (self *ComboBox) Popdown() {
	C.gtk_combo_box_popdown(self.object)
}

func (self *ComboBox) SetRowSeparatorFunc(f interface{}, data ...interface{}) {
	call, id := gobject.CreateCustomClosure(f, data...)
	_closures[id] = call
	C._gtk_combo_box_set_row_separator_func(self.object, C.gint64(id))
}

func (self *ComboBox) SetAddTearoffs(addTearoffs bool) {
	b := gobject.GBool(addTearoffs)
	defer b.Free()
	C.gtk_combo_box_set_add_tearoffs(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ComboBox) GetAddTearoffs() bool {
	b := C.gtk_combo_box_get_add_tearoffs(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ComboBox) SetTitle(title string) {
	s := gobject.GString(title)
	defer s.Free()
	C.gtk_combo_box_set_title(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *ComboBox) GetTitle() string {
	t := C.gtk_combo_box_get_title(self.object)
	s := C.g_strdup(t)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *ComboBox) SetFocusOnClick(focusOnClick bool) {
	b := gobject.GBool(focusOnClick)
	defer b.Free()
	C.gtk_combo_box_set_focus_on_click(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ComboBox) GetFocusOnClick() bool {
	b := C.gtk_combo_box_get_focus_on_click(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ComboBox) SetButtonSensitivity(gtk_Sensitivity int) {
	C.gtk_combo_box_set_button_sensitivity(self.object, C.GtkSensitivityType(gtk_Sensitivity))
}

func (self *ComboBox) GetButtonSensitivity() int {
	return int(C.gtk_combo_box_get_button_sensitivity(self.object))
}

func (self *ComboBox) GetHasEntry() bool {
	b := C.gtk_combo_box_get_has_entry(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ComboBox) SetEntryTextColumn(textColumn int) {
	C.gtk_combo_box_set_entry_text_column(self.object, C.gint(textColumn))
}

func (self *ComboBox) GetEntryTextColumn() int {
	return int(C.gtk_combo_box_get_entry_text_column(self.object))
}

func (self *ComboBox) SetPopupFixedWidth(fixed bool) {
	b := gobject.GBool(fixed)
	defer b.Free()
	C.gtk_combo_box_set_popup_fixed_width(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ComboBox) GetPopupFixedWidth() bool {
	b := C.gtk_combo_box_get_popup_fixed_width(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// End GtkComboBox
////////////////////////////// }}}

// GtkComboBoxText {{{
//////////////////////////////

// GtkComboBoxText type
type ComboBoxText struct {
	object *C.GtkComboBoxText
	*ComboBox
}

func NewComboBoxText() *ComboBoxText {
	ct := &ComboBoxText{}
	o := C.gtk_combo_box_text_new()
	ct.object = C.to_GtkComboBoxText(unsafe.Pointer(o))

	if gobject.IsObjectFloating(ct) {
		gobject.RefSink(ct)
	}
	ct.ComboBox = newComboBoxFromNative(unsafe.Pointer(o)).(*ComboBox)
	comboBoxTextFinalizer(ct)

	return ct
}

func NewComboBoxTextWithEntry() *ComboBoxText {
	ct := &ComboBoxText{}
	o := C.gtk_combo_box_text_new_with_entry()
	ct.object = C.to_GtkComboBoxText(unsafe.Pointer(o))

	if gobject.IsObjectFloating(ct) {
		gobject.RefSink(ct)
	}
	ct.ComboBox = newComboBoxFromNative(unsafe.Pointer(o)).(*ComboBox)
	comboBoxTextFinalizer(ct)

	return ct
}

// Clear ComboBoxText struct when it goes out of reach
func comboBoxTextFinalizer(ct *ComboBoxText) {
	runtime.SetFinalizer(ct, func(ct *ComboBoxText) { gobject.Unref(ct) })
}

// Conversion funcs
func newComboBoxTextFromNative(obj unsafe.Pointer) interface{} {
	ct := &ComboBoxText{}
	ct.object = C.to_GtkComboBoxText(obj)

	if gobject.IsObjectFloating(ct) {
		gobject.RefSink(ct)
	} else {
		gobject.Ref(ct)
	}
	ct.ComboBox = newComboBoxFromNative(obj).(*ComboBox)
	comboBoxTextFinalizer(ct)

	return ct
}

func nativeFromComboBoxText(ct interface{}) *gobject.GValue {
	combo, ok := ct.(*ComboBoxText)
	if ok {
		gv := gobject.CreateCGValue(GtkType.COMBO_BOX_TEXT, combo.ToNative())
		return gv
	}
	return nil
}

// To be Object-like
func (self ComboBoxText) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ComboBoxText) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ComboBoxText) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ComboBoxText) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Combo-like
func (self ComboBoxText) Combo() *ComboBox {
	return self.ComboBox
}

// ComboBoxText interface

func (self *ComboBoxText) Append(id, text string) {
	cid := gobject.GString(id)
	defer cid.Free()
	ctext := gobject.GString(text)
	defer ctext.Free()

	C.gtk_combo_box_text_append(self.object, (*C.gchar)(cid.GetPtr()), (*C.gchar)(ctext.GetPtr()))
}

func (self *ComboBoxText) Prepend(id, text string) {
	cid := gobject.GString(id)
	defer cid.Free()
	ctext := gobject.GString(text)
	defer ctext.Free()

	C.gtk_combo_box_text_prepend(self.object, (*C.gchar)(cid.GetPtr()), (*C.gchar)(ctext.GetPtr()))
}

func (self *ComboBoxText) Insert(position int, id, text string) {
	cid := gobject.GString(id)
	defer cid.Free()
	ctext := gobject.GString(text)
	defer ctext.Free()

	C.gtk_combo_box_text_insert(self.object, C.gint(position), (*C.gchar)(cid.GetPtr()), (*C.gchar)(ctext.GetPtr()))
}

func (self *ComboBoxText) AppendText(text string) {
	s := gobject.GString(text)
	defer s.Free()
	C.gtk_combo_box_text_append_text(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *ComboBoxText) PrependText(text string) {
	s := gobject.GString(text)
	defer s.Free()
	C.gtk_combo_box_text_prepend_text(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *ComboBoxText) InsertText(position int, text string) {
	s := gobject.GString(text)
	defer s.Free()
	C.gtk_combo_box_text_insert_text(self.object, C.gint(position), (*C.gchar)(s.GetPtr()))
}

func (self *ComboBoxText) Remove(position int) {
	C.gtk_combo_box_text_remove(self.object, C.gint(position))
}

func (self *ComboBoxText) RemoveAll() {
	C.gtk_combo_box_text_remove_all(self.object)
}

func (self *ComboBoxText) GetActiveText() string {
	s := C.gtk_combo_box_text_get_active_text(self.object)
	return gobject.GoString(unsafe.Pointer(s))
}

//////////////////////////////
// End GtkComboBoxText
////////////////////////////// }}}

// GtkMenuShell {{{
//////////////////////////////

// GtkMenuShell Type
type MenuShell struct {
	object *C.GtkMenuShell
	*Container
}

// Clear MenuShell struct when it goes out of reach
func menuShellFinalizer(m *MenuShell) {
	runtime.SetFinalizer(m, func(m *MenuShell) { gobject.Unref(m) })
}

// Conversion function for gobject registration map
func newMenuShellFromNative(obj unsafe.Pointer) interface{} {
	m := &MenuShell{}
	m.object = C.to_GtkMenuShell(obj)

	if gobject.IsObjectFloating(m) {
		gobject.RefSink(m)
	} else {
		gobject.Ref(m)
	}
	m.Container = NewContainer(obj)
	menuShellFinalizer(m)

	return m
}

func nativeFromMenuShell(m interface{}) *gobject.GValue {
	ms, ok := m.(*MenuShell)
	if ok {
		gv := gobject.CreateCGValue(GtkType.MENU_SHELL, ms.ToNative())
		return gv
	}
	return nil
}

// Conversion for GtkMenuDirection
func newGtkMenuDirectionFromNative(obj unsafe.Pointer) interface{} {
	return int(*((*C.gint)(obj)))
}

// To be Object-like
func (self MenuShell) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self MenuShell) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self MenuShell) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self MenuShell) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self MenuShell) C() *Container {
	return self.Container
}

// MenuShell interface

func (self *MenuShell) Append(child WidgetLike) {
	C.gtk_menu_shell_append(self.object, child.W().object)
}

func (self *MenuShell) Prepend(child WidgetLike) {
	C.gtk_menu_shell_prepend(self.object, child.W().object)
}

func (self *MenuShell) Insert(child WidgetLike, position int) {
	C.gtk_menu_shell_insert(self.object, child.W().object, C.gint(position))
}

func (self *MenuShell) Deactivate() {
	C.gtk_menu_shell_deactivate(self.object)
}

func (self *MenuShell) SelectItem(menuItem WidgetLike) {
	C.gtk_menu_shell_select_item(self.object, menuItem.W().object)
}

func (self *MenuShell) SelectFirst(searchSensitive bool) {
	b := gobject.GBool(searchSensitive)
	defer b.Free()

	C.gtk_menu_shell_select_first(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *MenuShell) Deselect() {
	C.gtk_menu_shell_deselect(self.object)
}

func (self *MenuShell) ActivateItem(menuItem WidgetLike, forceDeactivate bool) {
	b := gobject.GBool(forceDeactivate)
	defer b.Free()
	C.gtk_menu_shell_activate_item(self.object, menuItem.W().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *MenuShell) Cancel() {
	C.gtk_menu_shell_cancel(self.object)
}

func (self *MenuShell) SetTakeFocus(takeFocus bool) {
	b := gobject.GBool(takeFocus)
	defer b.Free()
	C.gtk_menu_shell_set_take_focus(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *MenuShell) GetTakeFocus() bool {
	b := C.gtk_menu_shell_get_take_focus(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *MenuShell) GetSelectedItem() WidgetLike {
	w := C.gtk_menu_shell_get_selected_item(self.object)

	if w == nil {
		return nil
	}

	if wid, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

func (self *MenuShell) GetParentShell() WidgetLike {
	w := C.gtk_menu_shell_get_parent_shell(self.object)

	if w == nil {
		return nil
	}

	if wid, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

//////////////////////////////
// END GtkMenuShell
////////////////////////////// }}}

// GtkMenu {{{
//////////////////////////////

// GtkMenu Type
type Menu struct {
	object *C.GtkMenu
	*MenuShell
}

func NewMenu() *Menu {
	m := &Menu{}
	o := C.gtk_menu_new()
	m.object = C.to_GtkMenu(unsafe.Pointer(o))

	if gobject.IsObjectFloating(m) {
		gobject.RefSink(m)
	}
	m.MenuShell = newMenuShellFromNative(unsafe.Pointer(o)).(*MenuShell)
	menuFinalizer(m)

	return m
}

// Clear Menu struct when it goes out of reach
func menuFinalizer(m *Menu) {
	runtime.SetFinalizer(m, func(m *Menu) { gobject.Unref(m) })
}

// Conversion function for gobject registration map
func newMenuFromNative(obj unsafe.Pointer) interface{} {
	m := &Menu{}
	m.object = C.to_GtkMenu(obj)

	if gobject.IsObjectFloating(m) {
		gobject.RefSink(m)
	} else {
		gobject.Ref(m)
	}
	m.MenuShell = newMenuShellFromNative(obj).(*MenuShell)
	menuFinalizer(m)

	return m
}

func nativeFromMenu(menu interface{}) *gobject.GValue {
	m, ok := menu.(*Menu)
	if ok {
		gv := gobject.CreateCGValue(GtkType.MENU_SHELL, m.ToNative())
		return gv
	}
	return nil
}

// To be Object-like
func (self Menu) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Menu) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Menu) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Menu) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be MenuShell-like
func (self Menu) MShell() *MenuShell {
	return self.MenuShell
}

// Menu interface

func (self *Menu) SetScreen(screen *gdk3.Screen) {
	if screen == nil {
		C.gtk_menu_set_screen(self.object, nil)
	}
	C.gtk_menu_set_screen(self.object, (*C.GdkScreen)(screen.ToNative()))
}

func (self *Menu) ReorderChild(child WidgetLike, position int) {
	C.gtk_menu_reorder_child(self.object, child.W().object, C.gint(position))
}

func (self *Menu) Attach(child WidgetLike, left, right, top, bottom uint) {
	C.gtk_menu_attach(self.object, child.W().object, C.guint(left),
		C.guint(right), C.guint(top), C.guint(bottom))
}

func (self *Menu) PopupForDevice(device *gdk3.Device, parentMenuShell WidgetLike, parentMenuItem WidgetLike,
	button uint, activateTime uint32, f interface{}, data ...interface{}) {

	var call gobject.ClosureFunc
	var id int64 = 0
	if f != nil {
		call, id = gobject.CreateCustomClosure(f, data...)
		_closures[id] = call
	}
	var pms, pmi *C.GtkWidget = nil, nil

	if parentMenuShell != nil {
		pms = parentMenuShell.W().object
	}

	if parentMenuItem != nil {
		pmi = parentMenuItem.W().object
	}

	C._gtk_menu_popup_for_device(self.object, (*C.GdkDevice)(device.ToNative()), pms,
		pmi, C.guint(button), C.guint32(activateTime), C.gint64(id))
}

func (self *Menu) Popup(parentMenuShell WidgetLike, parentMenuItem WidgetLike,
	button uint, activateTime uint32, f interface{}, data ...interface{}) {

	var call gobject.ClosureFunc
	var id int64 = 0
	if f != nil {
		call, id = gobject.CreateCustomClosure(f, data...)
		_closures[id] = call
	}
	var pms, pmi *C.GtkWidget = nil, nil

	if parentMenuShell != nil {
		pms = parentMenuShell.W().object
	}

	if parentMenuItem != nil {
		pmi = parentMenuItem.W().object
	}

	C._gtk_menu_popup(self.object, pms, pmi, C.guint(button), C.guint32(activateTime), C.gint64(id))

	if id != 0 {
		delete(_closures, id)
	}
}

func (self *Menu) SetAccelGroup(accelGroup *AccelGroup) {
	C.gtk_menu_set_accel_group(self.object, accelGroup.object)
}

func (self *Menu) GetAccelGroup() *AccelGroup {
	a := C.gtk_menu_get_accel_group(self.object)

	if a == nil {
		return nil
	}

	if acc, err := gobject.ConvertToGo(unsafe.Pointer(a)); err == nil {
		return acc.(*AccelGroup)
	}
	return nil
}

func (self *Menu) SetAccelPath(accelPath string) {
	s := gobject.GString(accelPath)
	defer s.Free()
	C.gtk_menu_set_accel_path(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Menu) GetAccelPath() string {
	s := C.gtk_menu_get_accel_path(self.object)
	return gobject.GoString(unsafe.Pointer(&s))
}

func (self *Menu) SetTitle(title interface{}) {
	if title == nil {
		C.gtk_menu_set_title(self.object, nil)
	}
	if t, ok := title.(string); ok {
		s := gobject.GString(t)
		defer s.Free()
		C.gtk_menu_set_title(self.object, (*C.gchar)(s.GetPtr()))
	}
}

func (self *Menu) GetTitle() string {
	s := C.gtk_menu_get_title(self.object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(C.g_strdup(s)))
}

func (self *Menu) SetMonitor(monitorNum int) {
	C.gtk_menu_set_monitor(self.object, C.gint(monitorNum))
}

func (self *Menu) GetMonitor() int {
	return int(C.gtk_menu_get_monitor(self.object))
}

func (self *Menu) GetTearoffState() bool {
	b := C.gtk_menu_get_tearoff_state(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Menu) SetReserverToggleSize(reserveToggleSize bool) {
	b := gobject.GBool(reserveToggleSize)
	defer b.Free()
	C.gtk_menu_set_reserve_toggle_size(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Menu) GetReserveToggleSize() bool {
	b := C.gtk_menu_get_reserve_toggle_size(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Menu) Popdown() {
	C.gtk_menu_popdown(self.object)
}

func (self *Menu) Reposition() {
	C.gtk_menu_reposition(self.object)
}

func (self *Menu) GetActive() WidgetLike {
	w := C.gtk_menu_get_active(self.object)
	if wid, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

func (self *Menu) SetActive(index uint) {
	C.gtk_menu_set_active(self.object, C.guint(index))
}

func (self *Menu) SetTearoffState(tornOff bool) {
	b := gobject.GBool(tornOff)
	defer b.Free()
	C.gtk_menu_set_tearoff_state(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Menu) AttachToWidget(attachWidget WidgetLike, f interface{}, data ...interface{}) {
	call, id := gobject.CreateCustomClosure(f, data...)
	_closures[id] = call

	C._gtk_menu_attach_to_widget(self.object, attachWidget.W().object, C.gint64(id))
}

func (self *Menu) Detach() {
	C.gtk_menu_detach(self.object)
}

func (self *Menu) GetAttachWidget() WidgetLike {
	w := C.gtk_menu_get_attach_widget(self.object)

	if w != nil {
		if wid, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
			return wid.(WidgetLike)
		}
	}
	return nil
}

func (self *Menu) GetForAttachWidget(w WidgetLike) *glib.GList {
	clist := C.gtk_menu_get_for_attach_widget(w.W().object)

	goList := glib.NewGListFromNative(unsafe.Pointer(clist))
	goList.ConversionFunc = func(obj unsafe.Pointer) interface{} {
		if w, err := gobject.ConvertToGo(obj); err == nil {
			return w.(WidgetLike)
		}
		return nil
	}

	return goList
}

//////////////////////////////
// END GtkMenu
////////////////////////////// }}}

// GtkMenuBar {{{
//////////////////////////////

// GtkMenuBar Type
type MenuBar struct {
	object *C.GtkMenuBar
	*MenuShell
}

func NewMenuBar() *MenuBar {
	m := &MenuBar{}
	o := C.gtk_menu_bar_new()
	m.object = C.to_GtkMenuBar(unsafe.Pointer(o))

	if gobject.IsObjectFloating(m) {
		gobject.RefSink(m)
	}
	m.MenuShell = newMenuShellFromNative(unsafe.Pointer(o)).(*MenuShell)
	menuBarFinalizer(m)

	return m
}

// Clear MenuBar struct when it goes out of reach
func menuBarFinalizer(m *MenuBar) {
	runtime.SetFinalizer(m, func(m *MenuBar) { gobject.Unref(m) })
}

// Conversion function for gobject registration map
func newMenuBarFromNative(obj unsafe.Pointer) interface{} {
	m := &MenuBar{}
	m.object = C.to_GtkMenuBar(obj)

	if gobject.IsObjectFloating(m) {
		gobject.RefSink(m)
	} else {
		gobject.Ref(m)
	}
	m.MenuShell = newMenuShellFromNative(obj).(*MenuShell)
	menuBarFinalizer(m)

	return m
}

func nativeFromMenuBar(menu interface{}) *gobject.GValue {
	m, ok := menu.(*MenuBar)
	if ok {
		gv := gobject.CreateCGValue(GtkType.MENU_BAR, m.ToNative())
		return gv
	}
	return nil
}

// To be Object-like
func (self MenuBar) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self MenuBar) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self MenuBar) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self MenuBar) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be MenuShell-like
func (self MenuBar) MShell() *MenuShell {
	return self.MenuShell
}

// MenuBar interface

func (self *MenuBar) SetPackDirection(gtk_PackDirection int) {
	C.gtk_menu_bar_set_pack_direction(self.object, C.GtkPackDirection(gtk_PackDirection))
}

func (self *MenuBar) GetPackDirection() int {
	return int(C.gtk_menu_bar_get_pack_direction(self.object))
}

func (self *MenuBar) SetChildPackDirection(gtk_PackDirection int) {
	C.gtk_menu_bar_set_child_pack_direction(self.object, C.GtkPackDirection(gtk_PackDirection))
}

func (self *MenuBar) GetChildPackDirection() int {
	return int(C.gtk_menu_bar_get_child_pack_direction(self.object))
}

//////////////////////////////
// END GtkMenuBar
////////////////////////////// }}}

// GtkMenuItem {{{
//////////////////////////////

// MenuItem type
type MenuItem struct {
	object *C.GtkMenuItem
	*Bin
}

func NewMenuItem() *MenuItem {
	mi := &MenuItem{}
	o := C.gtk_menu_item_new()
	mi.object = C.to_GtkMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	}
	mi.Bin = NewBin(unsafe.Pointer(o))
	menuItemFinalizer(mi)

	return mi
}

func NewMenuItemWithLabel(label string) *MenuItem {
	mi := &MenuItem{}
	s := gobject.GString(label)
	defer s.Free()
	o := C.gtk_menu_item_new_with_label((*C.gchar)(s.GetPtr()))
	mi.object = C.to_GtkMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	}
	mi.Bin = NewBin(unsafe.Pointer(o))
	menuItemFinalizer(mi)

	return mi
}

func NewMenuItemWithMnemonic(label string) *MenuItem {
	mi := &MenuItem{}
	s := gobject.GString(label)
	defer s.Free()
	o := C.gtk_menu_item_new_with_mnemonic((*C.gchar)(s.GetPtr()))
	mi.object = C.to_GtkMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	}
	mi.Bin = NewBin(unsafe.Pointer(o))
	menuItemFinalizer(mi)

	return mi
}

// Clear MenuItem struct when it goes out of reach
func menuItemFinalizer(m *MenuItem) {
	runtime.SetFinalizer(m, func(m *MenuItem) { gobject.Unref(m) })
}

// Conversion function for gobject registration map
func newMenuItemFromNative(obj unsafe.Pointer) interface{} {
	mi := &MenuItem{}
	mi.object = C.to_GtkMenuItem(obj)

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	} else {
		gobject.Ref(mi)
	}
	mi.Bin = NewBin(obj)
	menuItemFinalizer(mi)

	return mi
}

func nativeFromMenuItem(m interface{}) *gobject.GValue {
	mi, ok := m.(*MenuItem)
	if ok {
		gv := gobject.CreateCGValue(GtkType.MENU_ITEM, mi.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self MenuItem) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self MenuItem) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self MenuItem) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self MenuItem) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be bin-like
func (self MenuItem) CBin() *Bin {
	return self.Bin
}

// MenuItem interface

func (self *MenuItem) GetLabel() string {
	s := C.gtk_menu_item_get_label(self.object)

	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}

	return ""
}

func (self *MenuItem) SetLabel(label string) {
	s := gobject.GString(label)
	defer s.Free()
	C.gtk_menu_item_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *MenuItem) GetUseUnderline() bool {
	b := C.gtk_menu_item_get_use_underline(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *MenuItem) SetUseUnderline(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_menu_item_set_use_underline(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *MenuItem) SetSubmenu(submenu WidgetLike) {
	if submenu == nil {
		C.gtk_menu_item_set_submenu(self.object, nil)
		return
	}
	C.gtk_menu_item_set_submenu(self.object, submenu.W().object)
}

func (self *MenuItem) GetSubmenu() WidgetLike {
	w := C.gtk_menu_item_get_submenu(self.object)

	if w != nil {
		if sub, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
			return sub.(WidgetLike)
		}
	}
	return nil
}

func (self *MenuItem) SetAccelPath(accelPath string) {
	s := gobject.GString(accelPath)
	defer s.Free()
	C.gtk_menu_item_set_accel_path(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *MenuItem) GetAccelPath() string {
	s := C.gtk_menu_item_get_accel_path(self.object)

	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}

func (self *MenuItem) Select() {
	C.gtk_menu_item_select(self.object)
}

func (self *MenuItem) Deselect() {
	C.gtk_menu_item_deselect(self.object)
}

func (self *MenuItem) Activate() {
	C.gtk_menu_item_activate(self.object)
}

func (self *MenuItem) ToggleSizeRequest(requisition int) {
	var r C.gint
	r = C.gint(requisition)
	C.gtk_menu_item_toggle_size_request(self.object, &r)
}

func (self *MenuItem) ToggleSizeAllocate(allocation int) {
	var a C.gint
	a = C.gint(allocation)
	C.gtk_menu_item_toggle_size_request(self.object, &a)
}

func (self *MenuItem) GetReserveIndicator() bool {
	b := C.gtk_menu_item_get_reserve_indicator(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *MenuItem) SetReserveIndicator(reserve bool) {
	b := gobject.GBool(reserve)
	defer b.Free()
	C.gtk_menu_item_set_reserve_indicator(self.object, *((*C.gboolean)(b.GetPtr())))
}

//////////////////////////////
// END GtkMenuItem
////////////////////////////// }}}

// GtkCheckMenuItem {{{
//////////////////////////////

// CheckMenuItem type
type CheckMenuItem struct {
	object *C.GtkCheckMenuItem
	*MenuItem
}

func NewCheckMenuItem() *CheckMenuItem {
	mi := &CheckMenuItem{}
	o := C.gtk_check_menu_item_new()
	mi.object = C.to_GtkCheckMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	}
	mi.MenuItem = newMenuItemFromNative(unsafe.Pointer(o)).(*MenuItem)
	checkMenuItemFinalizer(mi)

	return mi
}

func NewCheckMenuItemWithLabel(label string) *CheckMenuItem {
	mi := &CheckMenuItem{}
	s := gobject.GString(label)
	defer s.Free()
	o := C.gtk_check_menu_item_new_with_label((*C.gchar)(s.GetPtr()))
	mi.object = C.to_GtkCheckMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	}
	mi.MenuItem = newMenuItemFromNative(unsafe.Pointer(o)).(*MenuItem)
	checkMenuItemFinalizer(mi)

	return mi
}

func NewCheckMenuItemWithMnemonic(label string) *CheckMenuItem {
	mi := &CheckMenuItem{}
	s := gobject.GString(label)
	defer s.Free()
	o := C.gtk_check_menu_item_new_with_mnemonic((*C.gchar)(s.GetPtr()))
	mi.object = C.to_GtkCheckMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	}
	mi.MenuItem = newMenuItemFromNative(unsafe.Pointer(o)).(*MenuItem)
	checkMenuItemFinalizer(mi)

	return mi
}

// Clear CheckMenuItem struct when it goes out of reach
func checkMenuItemFinalizer(m *CheckMenuItem) {
	runtime.SetFinalizer(m, func(m *CheckMenuItem) { gobject.Unref(m) })
}

// Conversion function for gobject registration map
func newCheckMenuItemFromNative(obj unsafe.Pointer) interface{} {
	cm := &CheckMenuItem{}
	cm.object = C.to_GtkCheckMenuItem(obj)

	if gobject.IsObjectFloating(cm) {
		gobject.RefSink(cm)
	} else {
		gobject.Ref(cm)
	}
	cm.MenuItem = newMenuItemFromNative(obj).(*MenuItem)
	checkMenuItemFinalizer(cm)

	return cm
}

func nativeFromCheckMenuItem(m interface{}) *gobject.GValue {
	cm, ok := m.(*CheckMenuItem)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CHECK_MENU_ITEM, cm.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CheckMenuItem) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CheckMenuItem) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CheckMenuItem) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CheckMenuItem) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be MenuItem-like
func (self CheckMenuItem) MItem() *MenuItem {
	return self.MenuItem
}

// CheckMenuItem interface

func (self *CheckMenuItem) GetActive() bool {
	b := C.gtk_check_menu_item_get_active(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CheckMenuItem) SetActive(isActive bool) {
	b := gobject.GBool(isActive)
	defer b.Free()
	C.gtk_check_menu_item_set_active(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CheckMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled(self.object)
}

func (self *CheckMenuItem) GetInconsistent() bool {
	b := C.gtk_check_menu_item_get_inconsistent(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CheckMenuItem) SetInconsistent(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_check_menu_item_set_inconsistent(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CheckMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	b := gobject.GBool(drawAsRadio)
	defer b.Free()
	C.gtk_check_menu_item_set_draw_as_radio(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CheckMenuItem) GetDrawAsRadio() bool {
	b := C.gtk_check_menu_item_get_draw_as_radio(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// END GtkCheckMenuItem
////////////////////////////// }}}

// GtkRadioMenuItem {{{
//////////////////////////////

// RadioMenuItem type
type RadioMenuItem struct {
	object *C.GtkRadioMenuItem
	*CheckMenuItem
}

func NewRadioMenuItem(group *glib.GSList) *RadioMenuItem {
	mr := &RadioMenuItem{}
	var list *C.GSList = nil
	if group != nil {
		list = (*C.GSList)(group.ToNative())
	}
	o := C.gtk_radio_menu_item_new(list)
	mr.object = C.to_GtkRadioMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mr) {
		gobject.RefSink(mr)
	}
	mr.CheckMenuItem = newCheckMenuItemFromNative(unsafe.Pointer(o)).(*CheckMenuItem)
	radioMenuItemFinalizer(mr)

	return mr
}

func NewRadioMenuItemWithLabel(group *glib.GSList, label string) *RadioMenuItem {
	mr := &RadioMenuItem{}
	var list *C.GSList = nil
	if group != nil {
		list = (*C.GSList)(group.ToNative())
	}
	s := gobject.GString(label)
	defer s.Free()
	o := C.gtk_radio_menu_item_new_with_label(list, (*C.gchar)(s.GetPtr()))
	mr.object = C.to_GtkRadioMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mr) {
		gobject.RefSink(mr)
	}
	mr.CheckMenuItem = newCheckMenuItemFromNative(unsafe.Pointer(o)).(*CheckMenuItem)
	radioMenuItemFinalizer(mr)

	return mr
}

func NewRadioMenuItemWithMnemonic(group *glib.GSList, label string) *RadioMenuItem {
	mr := &RadioMenuItem{}
	var list *C.GSList = nil
	if group != nil {
		list = (*C.GSList)(group.ToNative())
	}
	s := gobject.GString(label)
	defer s.Free()
	o := C.gtk_radio_menu_item_new_with_label(list, (*C.gchar)(s.GetPtr()))
	mr.object = C.to_GtkRadioMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mr) {
		gobject.RefSink(mr)
	}
	mr.CheckMenuItem = newCheckMenuItemFromNative(unsafe.Pointer(o)).(*CheckMenuItem)
	radioMenuItemFinalizer(mr)

	return mr
}

// Clear RadioMenuItem struct when it goes out of reach
func radioMenuItemFinalizer(m *RadioMenuItem) {
	runtime.SetFinalizer(m, func(m *RadioMenuItem) { gobject.Unref(m) })
}

// Conversion funcs
func newRadioMenuItemFromNative(obj unsafe.Pointer) interface{} {
	mr := &RadioMenuItem{}
	mr.object = C.to_GtkRadioMenuItem(obj)

	if gobject.IsObjectFloating(mr) {
		gobject.RefSink(mr)
	} else {
		gobject.Ref(mr)
	}
	mr.CheckMenuItem = newCheckMenuItemFromNative(obj).(*CheckMenuItem)
	radioMenuItemFinalizer(mr)

	return mr
}

func nativeFromRadioMenuItem(m interface{}) *gobject.GValue {
	mr, ok := m.(*RadioMenuItem)
	if ok {
		gv := gobject.CreateCGValue(GtkType.RADIO_MENU_ITEM, mr.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self RadioMenuItem) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self RadioMenuItem) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self RadioMenuItem) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self RadioMenuItem) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// RadioMenuItem interface

func (self *RadioMenuItem) GetGroup() *glib.GSList {
	l := C.gtk_radio_menu_item_get_group(self.object)

	if l != nil {
		goList := glib.NewGSListFromNative(unsafe.Pointer(l))
		goList.GC_Free = false
		goList.GC_FreeFull = false
		goList.ConversionFunc = newRadioMenuItemFromNative
		return goList
	}
	return nil
}

//////////////////////////////
// END GtkRadioMenuItem
////////////////////////////// }}}

// GtkImageMenuItem {{{
//////////////////////////////

// ImageMenuItem type
type ImageMenuItem struct {
	object *C.GtkImageMenuItem
	*MenuItem
}

func NewImageMenuItem() *ImageMenuItem {
	mi := &ImageMenuItem{}
	o := C.gtk_image_menu_item_new()
	mi.object = C.to_GtkImageMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	}
	mi.MenuItem = newMenuItemFromNative(unsafe.Pointer(o)).(*MenuItem)
	imageMenuItemFinalizer(mi)

	return mi
}

func NewImageMenuItemFromStock(stockId string, accelGroup *AccelGroup) *ImageMenuItem {
	mi := &ImageMenuItem{}
	s := gobject.GString(stockId)
	defer s.Free()

	var cal *C.GtkAccelGroup = nil
	if accelGroup != nil {
		cal = accelGroup.object
	}
	o := C.gtk_image_menu_item_new_from_stock((*C.gchar)(s.GetPtr()), cal)
	mi.object = C.to_GtkImageMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	}
	mi.MenuItem = newMenuItemFromNative(unsafe.Pointer(o)).(*MenuItem)
	imageMenuItemFinalizer(mi)

	return mi
}

func NewImageMenuItemWithLabel(label string) *ImageMenuItem {
	mi := &ImageMenuItem{}
	s := gobject.GString(label)
	defer s.Free()
	o := C.gtk_image_menu_item_new_with_label((*C.gchar)(s.GetPtr()))
	mi.object = C.to_GtkImageMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	}
	mi.MenuItem = newMenuItemFromNative(unsafe.Pointer(o)).(*MenuItem)
	imageMenuItemFinalizer(mi)

	return mi
}

func NewImageMenuItemWithMnemonic(label string) *ImageMenuItem {
	mi := &ImageMenuItem{}
	s := gobject.GString(label)
	defer s.Free()
	o := C.gtk_image_menu_item_new_with_mnemonic((*C.gchar)(s.GetPtr()))
	mi.object = C.to_GtkImageMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mi) {
		gobject.RefSink(mi)
	}
	mi.MenuItem = newMenuItemFromNative(unsafe.Pointer(o)).(*MenuItem)
	imageMenuItemFinalizer(mi)

	return mi
}

// Clear ImageMenuItem struct when it goes out of reach
func imageMenuItemFinalizer(m *ImageMenuItem) {
	runtime.SetFinalizer(m, func(m *ImageMenuItem) { gobject.Unref(m) })
}

// Conversion function for gobject registration map
func newImageMenuItemFromNative(obj unsafe.Pointer) interface{} {
	m := &ImageMenuItem{}
	m.object = C.to_GtkImageMenuItem(obj)

	if gobject.IsObjectFloating(m) {
		gobject.RefSink(m)
	} else {
		gobject.Ref(m)
	}
	m.MenuItem = newMenuItemFromNative(obj).(*MenuItem)
	imageMenuItemFinalizer(m)

	return m
}

func nativeFromImageMenuItem(m interface{}) *gobject.GValue {
	im, ok := m.(*CheckMenuItem)
	if ok {
		gv := gobject.CreateCGValue(GtkType.IMAGE_MENU_ITEM, im.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self ImageMenuItem) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ImageMenuItem) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ImageMenuItem) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ImageMenuItem) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be MenuItem-like
func (self ImageMenuItem) MItem() *MenuItem {
	return self.MenuItem
}

// ImageMenuItem interface

func (self *ImageMenuItem) SetImage(image WidgetLike) {
	if image == nil {
		return
	}
	C.gtk_image_menu_item_set_image(self.object, image.W().object)
}

func (self *ImageMenuItem) GetImage() WidgetLike {
	i := C.gtk_image_menu_item_get_image(self.object)

	if i != nil {
		if im, err := gobject.ConvertToGo(unsafe.Pointer(i)); err == nil {
			return im.(WidgetLike)
		}
	}
	return nil
}

func (self *ImageMenuItem) GetUseStock() bool {
	b := C.gtk_image_menu_item_get_use_stock(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ImageMenuItem) SetUseStock(useStock bool) {
	b := gobject.GBool(useStock)
	defer b.Free()
	C.gtk_image_menu_item_set_use_stock(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ImageMenuItem) GetAlwaysShowImage() bool {
	b := C.gtk_image_menu_item_get_always_show_image(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ImageMenuItem) SetAlwaysShowImage(alwaysShow bool) {
	b := gobject.GBool(alwaysShow)
	defer b.Free()
	C.gtk_image_menu_item_set_always_show_image(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ImageMenuItem) SetAccelGroup(accelGroup *AccelGroup) {
	C.gtk_image_menu_item_set_accel_group(self.object, accelGroup.object)
}

//////////////////////////////
// END GtkImageMenuItem
////////////////////////////// }}}

// GtkSeparatorMenuItem {{{
//////////////////////////////

// SeparatorMenuItem type
type SeparatorMenuItem struct {
	object *C.GtkSeparatorMenuItem
	*MenuItem
}

func NewSeparatorMenuItem() *SeparatorMenuItem {
	ms := &SeparatorMenuItem{}
	o := C.gtk_separator_menu_item_new()
	ms.object = C.to_GtkSeparatorMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(ms) {
		gobject.RefSink(ms)
	}
	ms.MenuItem = newMenuItemFromNative(unsafe.Pointer(o)).(*MenuItem)
	separatorMenuItemFinalizer(ms)

	return ms
}

// Clear SeparatorMenuItem struct when it goes out of reach
func separatorMenuItemFinalizer(m *SeparatorMenuItem) {
	runtime.SetFinalizer(m, func(m *SeparatorMenuItem) { gobject.Unref(m) })
}

// Conversion function for gobject registration map
func newSeparatorMenuItemFromNative(obj unsafe.Pointer) interface{} {
	sm := &SeparatorMenuItem{}
	sm.object = C.to_GtkSeparatorMenuItem(obj)

	if gobject.IsObjectFloating(sm) {
		gobject.RefSink(sm)
	} else {
		gobject.Ref(sm)
	}
	sm.MenuItem = newMenuItemFromNative(obj).(*MenuItem)
	separatorMenuItemFinalizer(sm)

	return sm
}

func nativeFromSeparatorMenuItem(m interface{}) *gobject.GValue {
	sm, ok := m.(*SeparatorMenuItem)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SEPARATOR_MENU_ITEM, sm.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self SeparatorMenuItem) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self SeparatorMenuItem) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self SeparatorMenuItem) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self SeparatorMenuItem) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be MenuItem-like
func (self SeparatorMenuItem) MItem() *MenuItem {
	return self.MenuItem
}

//////////////////////////////
// END GtkSeparatorMenuItem
////////////////////////////// }}}

// GtkTearoffMenuItem {{{
//////////////////////////////

// TearoffMenuItem type
type TearoffMenuItem struct {
	object *C.GtkTearoffMenuItem
	*MenuItem
}

func NewTearoffMenuItem() *TearoffMenuItem {
	mt := &TearoffMenuItem{}
	o := C.gtk_tearoff_menu_item_new()
	mt.object = C.to_GtkTearoffMenuItem(unsafe.Pointer(o))

	if gobject.IsObjectFloating(mt) {
		gobject.RefSink(mt)
	}
	mt.MenuItem = newMenuItemFromNative(unsafe.Pointer(o)).(*MenuItem)
	tearoffMenuItemFinalizer(mt)

	return mt
}

// Clear TearoffMenuItem struct when it goes out of reach
func tearoffMenuItemFinalizer(m *TearoffMenuItem) {
	runtime.SetFinalizer(m, func(m *TearoffMenuItem) { gobject.Unref(m) })
}

// Conversion function for gobject registration map
func newTearoffMenuItemFromNative(obj unsafe.Pointer) interface{} {
	mt := &TearoffMenuItem{}
	mt.object = C.to_GtkTearoffMenuItem(obj)

	if gobject.IsObjectFloating(mt) {
		gobject.RefSink(mt)
	} else {
		gobject.Ref(mt)
	}
	mt.MenuItem = newMenuItemFromNative(obj).(*MenuItem)
	tearoffMenuItemFinalizer(mt)

	return mt
}

func nativeFromTearoffMenuItem(m interface{}) *gobject.GValue {
	mt, ok := m.(*TearoffMenuItem)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TEAROFF_MENU_ITEM, mt.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TearoffMenuItem) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TearoffMenuItem) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TearoffMenuItem) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TearoffMenuItem) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be MenuItem-like
func (self TearoffMenuItem) MItem() *MenuItem {
	return self.MenuItem
}

//////////////////////////////
// END GtkTearoffMenuItem
////////////////////////////// }}}

// End Menus, Combo Box, Toolbar }}}

// Selectors (Color/File/Font) {{{

// GtkColorButton {{{
//////////////////////////////

// GtkColorButton type
type ColorButton struct {
	object *C.GtkColorButton
	*Button
}

// Create and return new color button
func NewColorButton() *ColorButton {
	b := &ColorButton{}
	o := C.gtk_color_button_new()
	b.object = C.to_GtkColorButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	colorButtonFinalizer(b)

	return b
}

func NewColorButtonWithColor(color gdk3.Color) *ColorButton {
    b := &ColorButton{}
    o := C.gtk_color_button_new_with_color((*C.GdkColor)(color.ToNative()))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	colorButtonFinalizer(b)

	return b
}

func NewColorButtonWithRGBA(rgba gdk3.RGBA) *ColorButton {
    b := &ColorButton{}
    o := C.gtk_color_button_new_with_rgba((*C.GdkRGBA)(rgba.ToNative()))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	colorButtonFinalizer(b)

	return b
}

// Clear ColorButton when it goes out of reach
func colorButtonFinalizer(b *ColorButton) {
	runtime.SetFinalizer(b, func(b *ColorButton) { gobject.Unref(b) })
}

// Conversion function for gobject registration map
func newColorButtonFromNative(obj unsafe.Pointer) interface{} {
	b := &ColorButton{}
	b.object = C.to_GtkColorButton(obj)

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	} else {
		gobject.Ref(b)
	}
	b.Button = newButtonFromNative(obj).(*Button)
	colorButtonFinalizer(b)

	return b
}

func nativeFromColorButton(b interface{}) *gobject.GValue {
	if but, ok := b.(*ColorButton); ok {
		gv := gobject.CreateCGValue(GtkType.COLOR_BUTTON, but.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self ColorButton) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ColorButton) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ColorButton) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ColorButton) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be button-like
func (self ColorButton) B() *Button {
	return self.Button
}

// ColorButton interface

func (self *ColorButton) SetColor(color gdk3.Color) {
    C.gtk_color_button_set_color(self.object, (*C.GdkColor)(color.ToNative()))
}

func (self *ColorButton) GetColor() gdk3.Color {
    var cl C.GdkColor
    C.gtk_color_button_get_color(self.object, &cl)

    if gdkColor, err := gobject.ConvertToGo(unsafe.Pointer(&cl), gdk3.GdkType.COLOR); err == nil {
        return gdkColor.(gdk3.Color)
    }

    return gdk3.Color{}
}

func (self *ColorButton) SetAlpha(alpha uint16) {
    C.gtk_color_button_set_alpha(self.object, C.guint16(alpha))
}

func (self *ColorButton) GetAlpha() uint16 {
    a := C.gtk_color_button_get_alpha(self.object)
    return uint16(a)
}

func (self *ColorButton) SetRGBA(rgba gdk3.RGBA) {
    C.gtk_color_button_set_rgba(self.object, (*C.GdkRGBA)(rgba.ToNative()))
}

func (self *ColorButton) GetRGBA() gdk3.RGBA {
    var cl C.GdkRGBA
    C.gtk_color_button_get_rgba(self.object, &cl)

    if gdkRGBA, err := gobject.ConvertToGo(unsafe.Pointer(&cl), gdk3.GdkType.RGBA); err == nil {
        return gdkRGBA.(gdk3.RGBA)
    }

    return gdk3.RGBA{}
}

func (self *ColorButton) SetUseAlpha(useAlpha bool) {
    b := gobject.GBool(useAlpha)
    defer b.Free()
    C.gtk_color_button_set_use_alpha(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ColorButton) GetUseAlpha() bool {
    b := C.gtk_color_button_get_use_alpha(self.object)
    return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ColorButton) SetTitle(title string) {
    s := gobject.GString(title)
    defer s.Free()
    C.gtk_color_button_set_title(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *ColorButton) GetTitle() string {
    s := C.gtk_color_button_get_title(self.object)
    return gobject.GoString(unsafe.Pointer(C.g_strdup(s)))
}


//////////////////////////////
// END GtkColorButton
////////////////////////////// }}}

// Selectors (Color/File/Font) }}}

// Action-based menus and toolbars {{{

// GtkAction {{{
//////////////////////////////

// GtkAction type
type Action struct {
	object *C.GtkAction
}

func NewAction(name string, label, tooltip, stockId interface{}) *Action {
	var cLabel, cTooltip, cStockId *C.gchar = nil, nil, nil

	if label != nil {
		if l, ok := label.(string); ok {
			lab := gobject.GString(l)
			defer lab.Free()
			cLabel = (*C.gchar)(lab.GetPtr())
		}
	}

	if tooltip != nil {
		if t, ok := tooltip.(string); ok {
			tt := gobject.GString(t)
			defer tt.Free()
			cTooltip = (*C.gchar)(tt.GetPtr())
		}
	}

	if stockId != nil {
		if s, ok := stockId.(string); ok {
			sid := gobject.GString(s)
			defer sid.Free()
			cStockId = (*C.gchar)(sid.GetPtr())
		}
	}

	cName := gobject.GString(name)
	defer cName.Free()

	a := &Action{}
	a.object = C.gtk_action_new((*C.gchar)(cName.GetPtr()), cLabel, cTooltip, cStockId)

	if gobject.IsObjectFloating(a) {
		gobject.RefSink(a)
	}
	actionFinalizer(a)

	return a
}

// Clear Action object when it goes of of reach
func actionFinalizer(a *Action) {
	runtime.SetFinalizer(a, func(a *Action) { gobject.Unref(a) })
}

// Conversion funcs
func newActionFromNative(obj unsafe.Pointer) interface{} {
	a := &Action{}
	a.object = C.to_GtkAction(obj)

	if gobject.IsObjectFloating(a) {
		gobject.RefSink(a)
	} else {
		gobject.Ref(a)
	}
	actionFinalizer(a)

	return a
}

func nativeFromAction(a interface{}) *gobject.GValue {
	ac, ok := a.(*Action)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ACTION, ac.ToNative())
		return gv
	}
	return nil
}

// To be Object-like
func (self Action) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Action) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Action) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Action) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Action interface

func (self *Action) GetName() string {
	s := C.gtk_action_get_name(self.object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(C.g_strdup(s)))
}

func (self *Action) IsSensitive() bool {
	b := C.gtk_action_is_sensitive(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Action) GetSensitive() bool {
	b := C.gtk_action_get_sensitive(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Action) SetSensitive(sensitive bool) {
	b := gobject.GBool(sensitive)
	defer b.Free()
	C.gtk_action_set_sensitive(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Action) IsVisible() bool {
	b := C.gtk_action_is_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Action) GetVisible() bool {
	b := C.gtk_action_get_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Action) SetVisible(visible bool) {
	b := gobject.GBool(visible)
	defer b.Free()
	C.gtk_action_set_visible(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Action) Activate() {
	C.gtk_action_activate(self.object)
}

func (self *Action) CreateIcon(gtk_IconSize int) WidgetLike {
	icon := C.gtk_action_create_icon(self.object, C.GtkIconSize(gtk_IconSize))

	if icon != nil {
		if i, err := gobject.ConvertToGo(unsafe.Pointer(icon)); err == nil {
			return i.(WidgetLike)
		}
	}
	return nil
}

func (self *Action) CreateMenuItem() WidgetLike {
	mi := C.gtk_action_create_menu_item(self.object)

	if mi != nil {
		if m, err := gobject.ConvertToGo(unsafe.Pointer(mi)); err == nil {
			return m.(WidgetLike)
		}
	}
	return nil
}

func (self *Action) CreateToolItem() WidgetLike {
	ti := C.gtk_action_create_tool_item(self.object)

	if ti != nil {
		if t, err := gobject.ConvertToGo(unsafe.Pointer(ti)); err == nil {
			return t.(WidgetLike)
		}
	}
	return nil
}

func (self *Action) CreateMenu() WidgetLike {
	m := C.gtk_action_create_menu(self.object)

	if m != nil {
		if menu, err := gobject.ConvertToGo(unsafe.Pointer(m)); err == nil {
			return menu.(WidgetLike)
		}
	}
	return nil
}

func (self *Action) GetProxies() *glib.GSList {
	l := C.gtk_action_get_proxies(self.object)

	if l != nil {
		goList := glib.NewGSListFromNative(unsafe.Pointer(l))
		goList.GC_Free = false
		goList.GC_FreeFull = false
		goList.ConversionFunc = func(obj unsafe.Pointer) interface{} {
			if w, err := gobject.ConvertToGo(obj); err == nil {
				return w
			}
			return nil
		}
		return goList
	}
	return nil
}

func (self *Action) ConnectAccelerator() {
	C.gtk_action_connect_accelerator(self.object)
}

func (self *Action) DisconnectAccelerator() {
	C.gtk_action_disconnect_accelerator(self.object)
}

func (self *Action) BlockActivate() {
	C.gtk_action_block_activate(self.object)
}

func (self *Action) UnblockActivate() {
	C.gtk_action_unblock_activate(self.object)
}

func (self *Action) GetAlwaysShowImage() bool {
	b := C.gtk_action_get_always_show_image(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Action) SetAlwaysShowImage(alwaysShow bool) {
	b := gobject.GBool(alwaysShow)
	defer b.Free()
	C.gtk_action_set_always_show_image(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Action) GetAccelPath() string {
	s := C.gtk_action_get_accel_path(self.object)
	if s != nil {
		return gobject.GoString(unsafe.Pointer(C.g_strdup(s)))
	}
	return ""
}

func (self *Action) SetAccelPath(accelPath string) {
	s := gobject.GString(accelPath)
	defer s.Free()
	C.gtk_action_set_accel_path(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Action) SetAccelGroup(accelGroup *AccelGroup) {
	if accelGroup == nil {
		C.gtk_action_set_accel_group(self.object, nil)
		return
	}

	C.gtk_action_set_accel_group(self.object, accelGroup.object)
}

func (self *Action) SetLabel(label string) {
	s := gobject.GString(label)
	defer s.Free()
	C.gtk_action_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Action) GetLabel() string {
	l := C.gtk_action_get_label(self.object)
	if l != nil {
		return gobject.GoString(unsafe.Pointer(l))
	}
	return ""
}

func (self *Action) SetShortLabel(shortLabel string) {
	s := gobject.GString(shortLabel)
	defer s.Free()
	C.gtk_action_set_short_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Action) GetShortLabel() string {
	l := C.gtk_action_get_short_label(self.object)
	if l != nil {
		return gobject.GoString(unsafe.Pointer(l))
	}
	return ""
}

func (self *Action) SetTooltip(tooltip string) {
	s := gobject.GString(tooltip)
	defer s.Free()
	C.gtk_action_set_tooltip(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Action) GetTooltip() string {
	t := C.gtk_action_get_tooltip(self.object)
	if t != nil {
		return gobject.GoString(unsafe.Pointer(t))
	}
	return ""
}

func (self *Action) SetStockId(stockId string) {
	s := gobject.GString(stockId)
	defer s.Free()
	C.gtk_action_set_stock_id(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Action) GetStockId() string {
	s := C.gtk_action_get_stock_id(self.object)
	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}

func (self *Action) SetIconName(iconName string) {
	s := gobject.GString(iconName)
	defer s.Free()
	C.gtk_action_set_icon_name(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Action) GetIconName() string {
	s := C.gtk_action_get_icon_name(self.object)
	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}

func (self *Action) SetVisibleHorizontal(visibleHorizontal bool) {
	b := gobject.GBool(visibleHorizontal)
	defer b.Free()
	C.gtk_action_set_visible_horizontal(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Action) GetVisibleHorizontal() bool {
	b := C.gtk_action_get_visible_horizontal(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Action) SetVisibleVertical(visibleVertical bool) {
	b := gobject.GBool(visibleVertical)
	defer b.Free()
	C.gtk_action_set_visible_vertical(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Action) GetVisibleVertical() bool {
	b := C.gtk_action_get_visible_vertical(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Action) SetIsImportant(isImportant bool) {
	b := gobject.GBool(isImportant)
	defer b.Free()
	C.gtk_action_set_is_important(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Action) GetIsImportant() bool {
	b := C.gtk_action_get_is_important(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//////////////////////////////
// End GtkAction
////////////////////////////// }}}

// End Action-based menus and toolbars }}}

// Layout Containers {{{

// GtkGrid {{{
//////////////////////////////

// GtkGrid Type
type Grid struct {
	object *C.GtkGrid
	*Container
	*Orientable
}

func NewGrid() *Grid {
	grid := &Grid{}
	o := C.gtk_grid_new()
	grid.object = C.to_GtkGrid(unsafe.Pointer(o))

	if gobject.IsObjectFloating(grid) {
		gobject.RefSink(grid)
	}
	grid.Container = NewContainer(unsafe.Pointer(o))
	grid.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	gridFinalizer(grid)

	return grid
}

// Clear Grid struct when it goes out of reach
func gridFinalizer(g *Grid) {
	runtime.SetFinalizer(g, func(g *Grid) { gobject.Unref(g) })
}

// Conversion function for gobject registration map
func newGridFromNative(obj unsafe.Pointer) interface{} {
	grid := &Grid{}
	grid.object = C.to_GtkGrid(obj)

	if gobject.IsObjectFloating(grid) {
		gobject.RefSink(grid)
	} else {
		gobject.Ref(grid)
	}
	grid.Container = NewContainer(obj)
	grid.Orientable = newOrientableFromNative(obj).(*Orientable)
	gridFinalizer(grid)

	return grid
}

func nativeFromGrid(g interface{}) *gobject.GValue {
	grid, ok := g.(*Grid)
	if ok {
		gv := gobject.CreateCGValue(GtkType.GRID, grid.ToNative())
		return gv
	}

	return nil
}

// To be Object-like
func (self Grid) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Grid) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Grid) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Grid) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self Grid) C() *Container {
	return self.Container
}

// To be Orientable-like
func (self Grid) IOrientable() *Orientable {
	return self.Orientable
}

// Grid interface
func (self *Grid) Attach(w WidgetLike, left, top, width, height int) {
	C.gtk_grid_attach(self.object, w.W().object, C.gint(left), C.gint(top),
		C.gint(width), C.gint(height))
}

func (self *Grid) AttachNextTo(child WidgetLike, sibling WidgetLike, side, width, height int) {
	C.gtk_grid_attach_next_to(self.object, child.W().object, sibling.W().object, C.GtkPositionType(side),
		C.gint(width), C.gint(height))
}

func (self *Grid) GetChildAt(left, top int) WidgetLike {
	c := C._gtk_grid_get_child_at(self.object, C.gint(left), C.gint(top))
	if c == nil {
		return nil
	}

	child, err := gobject.ConvertToGo(unsafe.Pointer(c))
	if err == nil {
		return child.(WidgetLike)
	}

	return nil
}

func (self *Grid) InsertRow(position int) {
	C._gtk_grid_insert_row(self.object, C.gint(position))
}

func (self *Grid) InsertColumn(position int) {
	C._gtk_grid_insert_column(self.object, C.gint(position))
}

func (self *Grid) InsertNextTo(sibling WidgetLike, side int) {
	C._gtk_grid_insert_next_to(self.object, sibling.W().object, C.GtkPositionType(side))
}

func (self *Grid) SetRowHomogeneous(homogeneous bool) {
	b := gobject.GBool(homogeneous)
	defer b.Free()
	C.gtk_grid_set_row_homogeneous(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Grid) GetRowHomogeneous() bool {
	b := C.gtk_grid_get_row_homogeneous(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Grid) SetRowSpacing(spacing uint) {
	C.gtk_grid_set_row_spacing(self.object, C.guint(spacing))
}

func (self *Grid) GetRowSpacing() uint {
	return uint(C.gtk_grid_get_row_spacing(self.object))
}

func (self *Grid) SetColumnHomogeneous(homogeneous bool) {
	b := gobject.GBool(homogeneous)
	defer b.Free()
	C.gtk_grid_set_column_homogeneous(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Grid) GetColumnHomogeneous() bool {
	b := C.gtk_grid_get_column_homogeneous(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Grid) SetColumnSpacing(spacing uint) {
	C.gtk_grid_set_column_spacing(self.object, C.guint(spacing))
}

func (self *Grid) GetColumnSpacing() uint {
	return uint(C.gtk_grid_get_column_spacing(self.object))
}

//////////////////////////////
// END GtkGrid
////////////////////////////// }}}

// GtkBox {{{
//////////////////////////////

// Box type
type Box struct {
	object *C.GtkBox
	*Container
	*Orientable
}

func NewBox(orientation int, spacing int) *Box {
	box := &Box{}
	o := C.gtk_box_new(C.GtkOrientation(orientation), C.gint(spacing))
	box.object = C.to_GtkBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(box) {
		gobject.RefSink(box)
	}
	box.Container = NewContainer(unsafe.Pointer(o))
	box.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	boxFinalizer(box)

	return box
}

func NewHBox(spacing int) *Box {
	return NewBox(GtkOrientation.HORIZONTAL, spacing)
}

func NewVBox(spacing int) *Box {
	return NewBox(GtkOrientation.VERTICAL, spacing)
}

// Clear Box struct when it goes out of reach
func boxFinalizer(box *Box) {
	runtime.SetFinalizer(box, func(box *Box) { gobject.Unref(box) })
}

// Conversion function for gobject registration map
func newBoxFromNative(obj unsafe.Pointer) interface{} {
	box := &Box{}
	box.object = C.to_GtkBox(obj)

	if gobject.IsObjectFloating(box) {
		gobject.RefSink(box)
	} else {
		gobject.Ref(box)
	}
	box.Container = NewContainer(obj)
	box.Orientable = newOrientableFromNative(obj).(*Orientable)
	boxFinalizer(box)

	return box
}

func nativeFromBox(b interface{}) *gobject.GValue {
	box, ok := b.(*Box)
	if ok {
		gv := gobject.CreateCGValue(GtkType.BOX, box.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self Box) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Box) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Box) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Box) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self Box) C() *Container {
	return self.Container
}

// To be orientable-like
func (self Box) IOrientable() *Orientable {
	return self.Orientable
}

// Box interface
func (self *Box) PackStart(w WidgetLike, expand bool, fill bool, padding uint) {
	e := gobject.GBool(expand)
	f := gobject.GBool(fill)
	defer e.Free()
	defer f.Free()
	C.gtk_box_pack_start(self.object, w.W().object, *((*C.gboolean)(e.GetPtr())), *((*C.gboolean)(e.GetPtr())), C.guint(padding))
}

func (self *Box) PackEnd(w WidgetLike, expand bool, fill bool, padding uint) {
	e := gobject.GBool(expand)
	f := gobject.GBool(fill)
	defer e.Free()
	defer f.Free()
	C.gtk_box_pack_end(self.object, w.W().object, *((*C.gboolean)(e.GetPtr())), *((*C.gboolean)(e.GetPtr())), C.guint(padding))
}

func (self *Box) GetHomogeneous() bool {
	h := gobject.GetProperty(self, "homogeneous")
	defer h.Free()
	return gobject.GoBool(h.GetPtr())
}

func (self *Box) SetHomogeneous(homogeneous bool) {
	gobject.SetProperty(self, "homogeneous", homogeneous)
}

func (self *Box) GetSpacing() int {
	return int(C.gtk_box_get_spacing(self.object))
}

func (self *Box) SetSpacing(spacing int) {
	gobject.SetProperty(self, "spacing", spacing)
}

func (self *Box) ReorderChild(w WidgetLike, position int) {
	C.gtk_box_reorder_child(self.object, w.W().object, C.gint(position))
}

func (self *Box) QueryChildPacking(w WidgetLike) (expand bool, fill bool, padding uint, ptype int) {
	var e C.gboolean
	var f C.gboolean
	var p C.guint
	var t C.GtkPackType
	C.gtk_box_query_child_packing(self.object, w.W().object, &e, &f, &p, &t)

	return gobject.GoBool(unsafe.Pointer(&e)), gobject.GoBool(unsafe.Pointer(&f)), uint(p), int(t)
}

func (self *Box) SetChildPacking(w WidgetLike, expand bool, fill bool, padding int, ptype int) {
	e := gobject.GBool(expand)
	f := gobject.GBool(fill)
	defer e.Free()
	defer f.Free()
	C.gtk_box_set_child_packing(self.object, w.W().object, *((*C.gboolean)(e.GetPtr())), *((*C.gboolean)(f.GetPtr())),
		C.guint(padding), C.GtkPackType(ptype))
}

//////////////////////////////
// END GtkBox
////////////////////////////// }}}

// GtkButtonBox {{{
//////////////////////////////

// ButtonBox type
type ButtonBox struct {
	object *C.GtkButtonBox
	*Box
	*Orientable
}

func NewButtonBox(orientation int) *ButtonBox {
	bb := &ButtonBox{}
	o := C.gtk_button_box_new(C.GtkOrientation(orientation))
	bb.object = C.to_GtkButtonBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(bb) {
		gobject.RefSink(bb)
	}
	bb.Box = newBoxFromNative(unsafe.Pointer(o)).(*Box)
	bb.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	buttonBoxFinalizer(bb)

	return bb
}

// Clear ButtonBox structure when it goes out of reach
func buttonBoxFinalizer(bb *ButtonBox) {
	runtime.SetFinalizer(bb, func(bb *ButtonBox) { gobject.Unref(bb) })
}

// Conversion function for gobject registration map
func newButtonBoxFromNative(obj unsafe.Pointer) interface{} {
	bb := &ButtonBox{}
	bb.object = C.to_GtkButtonBox(obj)

	if gobject.IsObjectFloating(bb) {
		gobject.RefSink(bb)
	}
	bb.Box = newBoxFromNative(obj).(*Box)
	bb.Orientable = newOrientableFromNative(obj).(*Orientable)
	buttonBoxFinalizer(bb)

	return bb
}

func nativeFromButtonBox(b interface{}) *gobject.GValue {
	bbox, ok := b.(*ButtonBox)
	if ok {
		gv := gobject.CreateCGValue(GtkType.BUTTON_BOX, bbox.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self ButtonBox) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ButtonBox) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ButtonBox) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ButtonBox) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Box-like
func (self ButtonBox) CBox() *Box {
	return self.Box
}

// To be Orientable-lik
func (self ButtonBox) IOrientable() *Orientable {
	return self.Orientable
}

// ButtonBox Interface
func (self *ButtonBox) GetLayout() int {
	l := C.gtk_button_box_get_layout(self.object)
	return int(l)
}

func (self *ButtonBox) GetChildSecondary(w WidgetLike) bool {
	cb := C.gtk_button_box_get_child_secondary(self.object, w.W().object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ButtonBox) GetChildNonHomogeneous(w WidgetLike) bool {
	cb := C._gtk_button_box_get_child_non_homogeneous(self.object, w.W().object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ButtonBox) SetLayout(layoutStyle int) {
	C.gtk_button_box_set_layout(self.object, C.GtkButtonBoxStyle(layoutStyle))
}

func (self *ButtonBox) SetChildSecondary(w WidgetLike, isSecondary bool) {
	cb := gobject.GBool(isSecondary)
	defer cb.Free()
	C.gtk_button_box_set_child_secondary(self.object, w.W().object, *((*C.gboolean)(cb.GetPtr())))
}

func (self *ButtonBox) SetChildNonHomogeneous(w WidgetLike, nonHomogeneous bool) {
	cb := gobject.GBool(nonHomogeneous)
	defer cb.Free()
	C._gtk_button_box_set_child_non_homogeneous(self.object, w.W().object, *((*C.gboolean)(cb.GetPtr())))
}

//////////////////////////////
// END GtkButtonBox
////////////////////////////// }}}

// GtkPaned {{{
//////////////////////////////

// Paned type
type Paned struct {
	object *C.GtkPaned
	*Container
	*Orientable
}

func NewPaned(gtk_Orientation int) *Paned {
	p := &Paned{}
	o := C.gtk_paned_new(C.GtkOrientation(gtk_Orientation))
	p.object = C.to_GtkPaned(unsafe.Pointer(o))

	if gobject.IsObjectFloating(p) {
		gobject.RefSink(p)
	}
	p.Container = NewContainer(unsafe.Pointer(o))
	p.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	panedFinalizer(p)

	return p
}

func NewHPaned() *Paned {
	return NewPaned(GtkOrientation.HORIZONTAL)
}

func NewVPaned() *Paned {
	return NewPaned(GtkOrientation.VERTICAL)
}

// Clear Paned struct when it goes out of reach
func panedFinalizer(p *Paned) {
	runtime.SetFinalizer(p, func(p *Paned) { gobject.Unref(p) })
}

// Conversion funcs
func newPanedFromNative(obj unsafe.Pointer) interface{} {
	p := &Paned{}
	p.object = C.to_GtkPaned(obj)

	if gobject.IsObjectFloating(p) {
		gobject.RefSink(p)
	} else {
		gobject.Ref(p)
	}
	p.Container = NewContainer(obj)
	p.Orientable = newOrientableFromNative(obj).(*Orientable)
	panedFinalizer(p)

	return p
}

func nativeFromPaned(p interface{}) *gobject.GValue {
	paned, ok := p.(*Paned)
	if ok {
		gv := gobject.CreateCGValue(GtkType.PANED, paned.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self Paned) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Paned) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Paned) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Paned) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self Paned) C() *Container {
	return self.Container
}

// To be orientable-like
func (self Paned) IOrientable() *Orientable {
	return self.Orientable
}

// Paned interface

func (self *Paned) Add1(child WidgetLike) {
	C.gtk_paned_add1(self.object, child.W().object)
}

func (self *Paned) Add2(child WidgetLike) {
	C.gtk_paned_add2(self.object, child.W().object)
}

func (self *Paned) Pack1(child WidgetLike, resize, shrink bool) {
	r := gobject.GBool(resize)
	defer r.Free()
	s := gobject.GBool(shrink)
	defer s.Free()

	C.gtk_paned_pack1(self.object, child.W().object, *((*C.gboolean)(r.GetPtr())), *((*C.gboolean)(s.GetPtr())))
}

func (self *Paned) Pack2(child WidgetLike, resize, shrink bool) {
	r := gobject.GBool(resize)
	defer r.Free()
	s := gobject.GBool(shrink)
	defer s.Free()

	C.gtk_paned_pack2(self.object, child.W().object, *((*C.gboolean)(r.GetPtr())), *((*C.gboolean)(s.GetPtr())))
}

func (self *Paned) GetChild1() WidgetLike {
	w := C.gtk_paned_get_child1(self.object)
	if w == nil {
		return nil
	}

	if wid, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

func (self *Paned) GetChild2() WidgetLike {
	w := C.gtk_paned_get_child2(self.object)
	if w == nil {
		return nil
	}

	if wid, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

func (self *Paned) SetPosition(position int) {
	C.gtk_paned_set_position(self.object, C.gint(position))
}

func (self *Paned) GetPosition() int {
	return int(C.gtk_paned_get_position(self.object))
}

func (self *Paned) GetHandleWindow() *gdk3.Window {
	w := C.gtk_paned_get_handle_window(self.object)
	if w == nil {
		return nil
	}

	if win, err := gobject.ConvertToGo(unsafe.Pointer(w)); err == nil {
		return win.(*gdk3.Window)
	}
	return nil
}

//////////////////////////////
// END GtkPaned
////////////////////////////// }}}

// GtkNotebook {{{
//////////////////////////////

// GtkNotebook type
type Notebook struct {
	object *C.GtkNotebook
	*Container
}

func NewNotebook() *Notebook {
	n := &Notebook{}
	o := C.gtk_notebook_new()
	n.object = C.to_GtkNotebook(unsafe.Pointer(o))

	if gobject.IsObjectFloating(n) {
		gobject.RefSink(n)
	}
	n.Container = NewContainer(unsafe.Pointer(o))
	notebookFinalizer(n)

	return n
}

// Clear Notebook struct when it goes out of reach
func notebookFinalizer(n *Notebook) {
	runtime.SetFinalizer(n, func(n *Notebook) { gobject.Unref(n) })
}

// Conversion functions
func newNotebookFromNative(obj unsafe.Pointer) interface{} {
	n := &Notebook{}
	n.object = C.to_GtkNotebook(obj)

	if gobject.IsObjectFloating(n) {
		gobject.RefSink(n)
	} else {
		gobject.Ref(n)
	}
	n.Container = NewContainer(obj)
	notebookFinalizer(n)

	return n
}

func nativeFromNotebook(note interface{}) *gobject.GValue {
	n, ok := note.(*Notebook)
	if ok {
		gv := gobject.CreateCGValue(GtkType.NOTEBOOK, n.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Notebook) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Notebook) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Notebook) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Notebook) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self Notebook) C() *Container {
	return self.Container
}

// Notebook interface

func (self *Notebook) AppendPage(child, tabLabel WidgetLike) int {
	var tab *C.GtkWidget = nil
	if tabLabel != nil {
		tab = tabLabel.W().object
	}
	return int(C.gtk_notebook_append_page(self.object, child.W().object, tab))
}

func (self *Notebook) AppendPageMenu(child, tabLabel, menuLabel WidgetLike) int {
	var t, m *C.GtkWidget = nil, nil

	if tabLabel != nil {
		t = tabLabel.W().object
	}

	if menuLabel != nil {
		m = menuLabel.W().object
	}
	return int(C.gtk_notebook_append_page_menu(self.object, child.W().object, t, m))
}

func (self *Notebook) PrependPage(child, tabLabel WidgetLike) int {
	var t *C.GtkWidget
	if tabLabel != nil {
		t = tabLabel.W().object
	}
	return int(C.gtk_notebook_prepend_page(self.object, child.W().object, t))
}

func (self *Notebook) PrependPageMenu(child, tabLabel, menuLabel WidgetLike) int {
	var t, m *C.GtkWidget = nil, nil

	if tabLabel != nil {
		t = tabLabel.W().object
	}

	if menuLabel != nil {
		m = menuLabel.W().object
	}
	return int(C.gtk_notebook_prepend_page_menu(self.object, child.W().object, t, m))
}

func (self *Notebook) InsertPage(child, tabLabel WidgetLike, position int) int {
	var t *C.GtkWidget = nil

	if tabLabel != nil {
		t = tabLabel.W().object
	}
	return int(C.gtk_notebook_insert_page(self.object, child.W().object, t, C.gint(position)))
}

func (self *Notebook) InsertPageMenu(child, tabLabel, menuLabel WidgetLike, position int) int {
	var t, m *C.GtkWidget = nil, nil

	if tabLabel != nil {
		t = tabLabel.W().object
	}

	if menuLabel != nil {
		m = menuLabel.W().object
	}
	return int(C.gtk_notebook_insert_page_menu(self.object, child.W().object, t, m, C.gint(position)))
}

func (self *Notebook) RemovePage(pageNum int) {
	C.gtk_notebook_remove_page(self.object, C.gint(pageNum))
}

func (self *Notebook) PageNum(child WidgetLike) int {
	return int(C.gtk_notebook_page_num(self.object, child.W().object))
}

func (self *Notebook) NextPage() {
	C.gtk_notebook_next_page(self.object)
}

func (self *Notebook) PrevPage() {
	C.gtk_notebook_prev_page(self.object)
}

func (self *Notebook) ReorderChild(child WidgetLike, position int) {
	C.gtk_notebook_reorder_child(self.object, child.W().object, C.gint(position))
}

func (self *Notebook) SetTabPos(gtk_PositionType int) {
	C.gtk_notebook_set_tab_pos(self.object, C.GtkPositionType(gtk_PositionType))
}

func (self *Notebook) SetShowTabs(showTabs bool) {
	b := gobject.GBool(showTabs)
	defer b.Free()
	C.gtk_notebook_set_show_tabs(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Notebook) SetShowBorder(showBorder bool) {
	b := gobject.GBool(showBorder)
	defer b.Free()
	C.gtk_notebook_set_show_border(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Notebook) SetScrollable(scrollable bool) {
	b := gobject.GBool(scrollable)
	defer b.Free()
	C.gtk_notebook_set_scrollable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Notebook) PopupEnable() {
	C.gtk_notebook_popup_enable(self.object)
}

func (self *Notebook) PopupDisable() {
	C.gtk_notebook_popup_disable(self.object)
}

func (self *Notebook) GetCurrentPage() int {
	return int(C.gtk_notebook_get_current_page(self.object))
}

func (self *Notebook) GetMenuLabel(child WidgetLike) WidgetLike {
	ml := C.gtk_notebook_get_menu_label(self.object, child.W().object)
	if ml == nil {
		return nil
	}

	l, err := gobject.ConvertToGo(unsafe.Pointer(ml))
	if err == nil {
		return l.(WidgetLike)
	}
	return nil
}

func (self *Notebook) GetNthPage(pageNum int) WidgetLike {
	w := C.gtk_notebook_get_nth_page(self.object, C.gint(pageNum))
	if w == nil {
		return nil
	}

	child, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return child.(WidgetLike)
	}
	return nil
}

func (self *Notebook) GetNPages() int {
	return int(C.gtk_notebook_get_n_pages(self.object))
}

func (self *Notebook) GetTabLabel(child WidgetLike) WidgetLike {
	w := C.gtk_notebook_get_tab_label(self.object, child.W().object)
	if w == nil {
		return nil
	}

	widget, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return widget.(WidgetLike)
	}
	return nil
}

func (self *Notebook) SetMenuLabel(child, menuLabel WidgetLike) {
	var m *C.GtkWidget = nil
	if menuLabel != nil {
		m = menuLabel.W().object
	}
	C.gtk_notebook_set_menu_label(self.object, child.W().object, m)
}

func (self *Notebook) SetMenuLabelText(child WidgetLike, menuText string) {
	s := gobject.GString(menuText)
	defer s.Free()
	C.gtk_notebook_set_menu_label_text(self.object, child.W().object, (*C.gchar)(s.GetPtr()))
}

func (self *Notebook) SetTabLabel(child, tabLabel WidgetLike) {
	var l *C.GtkWidget = nil
	if tabLabel != nil {
		l = tabLabel.W().object
	}
	C.gtk_notebook_set_tab_label(self.object, child.W().object, l)
}

func (self *Notebook) SetTabLabelText(child WidgetLike, tabText string) {
	s := gobject.GString(tabText)
	defer s.Free()
	C.gtk_notebook_set_tab_label_text(self.object, child.W().object, (*C.gchar)(s.GetPtr()))
}

func (self *Notebook) SetTabReorderable(child WidgetLike, reorderable bool) {
	b := gobject.GBool(reorderable)
	defer b.Free()
	C.gtk_notebook_set_tab_reorderable(self.object, child.W().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Notebook) SetTabDetachable(child WidgetLike, detachable bool) {
	b := gobject.GBool(detachable)
	defer b.Free()
	C.gtk_notebook_set_tab_detachable(self.object, child.W().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Notebook) GetMenuLabelText(child WidgetLike) string {
	s := C.gtk_notebook_get_menu_label_text(self.object, child.W().object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Notebook) GetScrollable() bool {
	b := C.gtk_notebook_get_scrollable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Notebook) GetShowBorder() bool {
	b := C.gtk_notebook_get_show_border(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Notebook) GetShowTabs() bool {
	b := C.gtk_notebook_get_show_tabs(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Notebook) GetTabLabelText(child WidgetLike) string {
	s := C.gtk_notebook_get_tab_label_text(self.object, child.W().object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Notebook) GetTabPos() (gtk_PositionType int) {
	return int(C.gtk_notebook_get_tab_pos(self.object))
}

func (self *Notebook) GetTabReorderable(child WidgetLike) bool {
	b := C.gtk_notebook_get_tab_reorderable(self.object, child.W().object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Notebook) GetTabDetachable(child WidgetLike) bool {
	b := C.gtk_notebook_get_tab_detachable(self.object, child.W().object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Notebook) GetTabHborder() uint {
	return uint(C.gtk_notebook_get_tab_hborder(self.object))
}

func (self *Notebook) GetTabVborder() uint {
	return uint(C.gtk_notebook_get_tab_vborder(self.object))
}

func (self *Notebook) SetCurrentPage(pageNum int) {
	C.gtk_notebook_set_current_page(self.object, C.gint(pageNum))
}

func (self *Notebook) SetGroupName(groupName string) {
	if groupName == "" {
		C.gtk_notebook_set_group_name(self.object, nil)
	}
	s := gobject.GString(groupName)
	defer s.Free()

	C.gtk_notebook_set_group_name(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Notebook) GetGroupName() string {
	s := C.gtk_notebook_get_group_name(self.object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Notebook) SetActionWidget(w WidgetLike, gtk_Pack int) {
	C.gtk_notebook_set_action_widget(self.object, w.W().object, C.GtkPackType(gtk_Pack))
}

func (self *Notebook) GetActionWidget(gtk_Pack int) WidgetLike {
	w := C.gtk_notebook_get_action_widget(self.object, C.GtkPackType(gtk_Pack))
	if w == nil {
		return nil
	}
	widget, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return widget.(WidgetLike)
	}
	return nil
}

//////////////////////////////
// End GtkNotebook
////////////////////////////// }}}

// GtkOrientable interface {{{
//////////////////////////////

// GtkOrientable type
type Orientable struct {
	object *C.GtkOrientable
}

// Clear Orientable when it goes out of reach
func orientableFinalizer(o *Orientable) {
	runtime.SetFinalizer(o, func(o *Orientable) { gobject.Unref(o) })
}

// Conversion funcs
func newOrientableFromNative(obj unsafe.Pointer) interface{} {
	o := &Orientable{}
	o.object = C.to_GtkOrientable(obj)

	if gobject.IsObjectFloating(o) {
		gobject.RefSink(o)
	} else {
		gobject.Ref(o)
	}
	orientableFinalizer(o)

	return o
}

func nativeFromOrientable(o interface{}) *gobject.GValue {
	ori, ok := o.(*Orientable)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ORIENTABLE, ori.ToNative())
		return gv
	}
	return nil
}

// Simulate object-like
func (self Orientable) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Orientable) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Orientable) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Orientable) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Orientable interface

func (self *Orientable) GetOrientation() int {
	return int(C.gtk_orientable_get_orientation(self.object))
}

func (self *Orientable) SetOrientation(gtk_Orientation int) {
	C.gtk_orientable_set_orientation(self.object, C.GtkOrientation(gtk_Orientation))
}

//////////////////////////////
// End GtkOrientable interface
////////////////////////////// }}}

// End Layout Containers }}}

// Ornaments {{{

// GtkFrame {{{
//////////////////////////////

// Frame type
type Frame struct {
	object *C.GtkFrame
	*Bin
}

func NewFrame(label string) *Frame {
	f := &Frame{}

	s := gobject.GString(label)
	defer s.Free()

	o := C.gtk_frame_new((*C.gchar)(s.GetPtr()))
	f.object = C.to_GtkFrame(unsafe.Pointer(o))

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	}
	f.Bin = NewBin(unsafe.Pointer(o))
	frameFinalizer(f)

	return f
}

// Clear Frame struct when it goes out of reach
func frameFinalizer(f *Frame) {
	runtime.SetFinalizer(f, func(f *Frame) { gobject.Unref(f) })
}

// Conversion function for gobject registration map
func newFrameFromNative(obj unsafe.Pointer) interface{} {
	f := &Frame{}
	f.object = C.to_GtkFrame(obj)

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	} else {
		gobject.Ref(f)
	}
	f.Bin = NewBin(obj)
	frameFinalizer(f)

	return f
}

func nativeFromFrame(frame interface{}) *gobject.GValue {
	f, ok := frame.(*Frame)
	if ok {
		gv := gobject.CreateCGValue(GtkType.FRAME, f.ToNative())
		return gv
	}

	return nil
}

// To be object like
func (self Frame) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Frame) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Frame) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Frame) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be bin-like
func (self Frame) CBin() *Bin {
	return self.Bin
}

// Frame interface
func (self *Frame) SetLabel(label string) {
	s := gobject.GString(label)
	defer s.Free()
	C.gtk_frame_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Frame) SetLabelWidget(w WidgetLike) {
	C.gtk_frame_set_label_widget(self.object, w.W().object)
}

func (self *Frame) SetLabelAlign(xalign, yalign float32) {
	C.gtk_frame_set_label_align(self.object, C.gfloat(xalign), C.gfloat(yalign))
}

func (self *Frame) SetShadowType(gtk_shadow int) {
	C.gtk_frame_set_shadow_type(self.object, C.GtkShadowType(gtk_shadow))
}

func (self *Frame) GetLabel() string {
	l := C.gtk_frame_get_label(self.object)
	return gobject.GoString(unsafe.Pointer(l))
}

func (self *Frame) GetLabelAlign() (xalign, yalign float32) {
	var cax C.gfloat
	var cay C.gfloat
	C.gtk_frame_get_label_align(self.object, &cax, &cay)
	return float32(cax), float32(cay)
}

func (self *Frame) GetLabelWidget() WidgetLike {
	cw := C.gtk_frame_get_label_widget(self.object)
	w, err := gobject.ConvertToGo(unsafe.Pointer(cw))
	if err != nil {
		return w.(WidgetLike)
	}
	return nil
}

func (self *Frame) GetShadowType() int {
	return int(C.gtk_frame_get_shadow_type(self.object))
}

//////////////////////////////
// END GtkFrame
////////////////////////////// }}}

// GtkSeparator {{{
//////////////////////////////

// GtkSeparator type
type Separator struct {
	object *C.GtkSeparator
	*Widget
	*Orientable
}

func NewSeparator(orientation int) *Separator {
	sep := &Separator{}
	o := C.gtk_separator_new(C.GtkOrientation(orientation))
	sep.object = C.to_GtkSeparator(unsafe.Pointer(o))

	if gobject.IsObjectFloating(sep) {
		gobject.RefSink(sep)
	}
	sep.Widget = NewWidget(unsafe.Pointer(o))
	sep.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	separatorFinalizer(sep)

	return sep
}

func NewHSeparator() *Separator {
	return NewSeparator(GtkOrientation.HORIZONTAL)
}

func NewVSeparator() *Separator {
	return NewSeparator(GtkOrientation.VERTICAL)
}

// Clear Separator struct when it goes out of reach
func separatorFinalizer(s *Separator) {
	runtime.SetFinalizer(s, func(s *Separator) { gobject.Unref(s) })
}

// Conversion function for gobject registration map
func newSeparatorFromNative(obj unsafe.Pointer) interface{} {
	sep := &Separator{}
	sep.object = C.to_GtkSeparator(obj)

	if gobject.IsObjectFloating(sep) {
		gobject.RefSink(sep)
	} else {
		gobject.Ref(sep)
	}
	sep.Widget = NewWidget(obj)
	sep.Orientable = newOrientableFromNative(obj).(*Orientable)
	separatorFinalizer(sep)

	return sep
}

func nativeFromSeparator(s interface{}) *gobject.GValue {
	sep, ok := s.(*Separator)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SEPARATOR, sep.ToNative())
		return gv
	}

	return nil
}

// To be object like
func (self Separator) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Separator) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Separator) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self Separator) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Separator) W() *Widget {
	return self.Widget
}

// To be Orientable-like
func (self Separator) IOrientable() *Orientable {
	return self.Orientable
}

//////////////////////////////
// END GtkSeparator
////////////////////////////// }}}

// End Ornaments }}}

// Scrolling {{{

// GtkScrollbar {{{
//////////////////////////////

// GtkScrollbar type
type Scrollbar struct {
	object *C.GtkScrollbar
	*Range
	*Orientable
}

func NewScrollbar(gtk_orientation int, adjustment *Adjustment) *Scrollbar {
	sb := &Scrollbar{}
	o := C.gtk_scrollbar_new(C.GtkOrientation(gtk_orientation), adjustment.object)
	sb.object = C.to_GtkScrollbar(unsafe.Pointer(o))

	if gobject.IsObjectFloating(sb) {
		gobject.RefSink(sb)
	}
	sb.Range = newRangeFromNative(unsafe.Pointer(o)).(*Range)
	sb.Orientable = newOrientableFromNative(unsafe.Pointer(o)).(*Orientable)
	scrollbarFinalizer(sb)

	return sb
}

func NewHScrollbar(adjustment *Adjustment) *Scrollbar {
	return NewScrollbar(GtkOrientation.HORIZONTAL, adjustment)
}

func NewVScrollbar(adjustment *Adjustment) *Scrollbar {
	return NewScrollbar(GtkOrientation.VERTICAL, adjustment)
}

// Clear Scrollbar struct when it goes out of reach
func scrollbarFinalizer(sb *Scrollbar) {
	runtime.SetFinalizer(sb, func(sb *Scrollbar) { gobject.Unref(sb) })
}

// Conversions functions
func newScrollbarFromNative(sb unsafe.Pointer) interface{} {
	scrollbar := &Scrollbar{}
	scrollbar.object = C.to_GtkScrollbar(sb)

	if gobject.IsObjectFloating(scrollbar) {
		gobject.RefSink(scrollbar)
	} else {
		gobject.Ref(scrollbar)
	}
	scrollbar.Range = newRangeFromNative(sb).(*Range)
	scrollbar.Orientable = newOrientableFromNative(sb).(*Orientable)
	scrollbarFinalizer(scrollbar)

	return scrollbar
}

func nativeFromScrollbar(sb interface{}) *gobject.GValue {
	scrollbar, ok := sb.(*Scrollbar)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SCROLLBAR, scrollbar.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Scrollbar) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Scrollbar) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Scrollbar) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Scrollbar) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Range-like
func (self Scrollbar) R() *Range {
	return self.Range
}

// To be Orientable-like
func (self Scrollbar) IOrientable() *Orientable {
	return self.Orientable
}

//////////////////////////////
// END GtkScrollbar
////////////////////////////// }}}

// GtkScrolledWindow {{{
//////////////////////////////

// GtkScrolledWindow type
type ScrolledWindow struct {
	object *C.GtkScrolledWindow
	*Bin
}

func NewScrolledWindow(hadjustment, vadjustment *Adjustment) *ScrolledWindow {
	sw := &ScrolledWindow{}
	var ha, va *C.GtkAdjustment
	ha, va = nil, nil

	if hadjustment != nil {
		ha = hadjustment.object
	}

	if vadjustment != nil {
		va = vadjustment.object
	}

	o := C.gtk_scrolled_window_new(ha, va)
	sw.object = C.to_GtkScrolledWindow(unsafe.Pointer(o))

	if gobject.IsObjectFloating(sw) {
		gobject.RefSink(sw)
	}
	sw.Bin = NewBin(unsafe.Pointer(o))
	scrolledWindowFinalizer(sw)

	return sw
}

// Clear ScrolledWindow struct when it goes out of reach
func scrolledWindowFinalizer(sw *ScrolledWindow) {
	runtime.SetFinalizer(sw, func(sw *ScrolledWindow) { gobject.Unref(sw) })
}

// Conversion functions
func newScrolledWindowFromNative(sw unsafe.Pointer) interface{} {
	scrolled := &ScrolledWindow{}
	scrolled.object = C.to_GtkScrolledWindow(sw)

	if gobject.IsObjectFloating(scrolled) {
		gobject.RefSink(scrolled)
	} else {
		gobject.Ref(scrolled)
	}
	scrolled.Bin = NewBin(sw)
	scrolledWindowFinalizer(scrolled)

	return scrolled
}

func nativeFromScrolledWindow(sw interface{}) *gobject.GValue {
	scrolled, ok := sw.(*ScrolledWindow)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SCROLLED_WINDOW, scrolled.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self ScrolledWindow) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ScrolledWindow) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ScrolledWindow) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ScrolledWindow) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be bin-like
func (self ScrolledWindow) CBin() *Bin {
	return self.Bin
}

// ScrolledWindow interface
func (self *ScrolledWindow) GetHadjustment() *Adjustment {
	ad := C.gtk_scrolled_window_get_hadjustment(self.object)
	a, err := gobject.ConvertToGo(unsafe.Pointer(ad))
	if err == nil {
		return a.(*Adjustment)
	}
	return nil
}

func (self *ScrolledWindow) GetVadjustment() *Adjustment {
	ad := C.gtk_scrolled_window_get_vadjustment(self.object)
	a, err := gobject.ConvertToGo(unsafe.Pointer(ad))
	if err == nil {
		return a.(*Adjustment)
	}
	return nil
}

func (self *ScrolledWindow) GetHScrollbar() WidgetLike {
	ad := C.gtk_scrolled_window_get_hscrollbar(self.object)
	a, err := gobject.ConvertToGo(unsafe.Pointer(ad))
	if err == nil {
		return a.(WidgetLike)
	}
	return nil
}

func (self *ScrolledWindow) GetVScrollbar() WidgetLike {
	ad := C.gtk_scrolled_window_get_vscrollbar(self.object)
	a, err := gobject.ConvertToGo(unsafe.Pointer(ad))
	if err == nil {
		return a.(WidgetLike)
	}
	return nil
}

func (self *ScrolledWindow) SetPolicy(gtk_policy_horizontal, gtk_policy_vertical int) {
	C.gtk_scrolled_window_set_policy(self.object, C.GtkPolicyType(gtk_policy_horizontal), C.GtkPolicyType(gtk_policy_vertical))
}

func (self *ScrolledWindow) AddWithViewport(child WidgetLike) {
	C.gtk_scrolled_window_add_with_viewport(self.object, child.W().object)
}

func (self *ScrolledWindow) SetPlacement(gtk_corner_window_placement int) {
	C.gtk_scrolled_window_set_placement(self.object, C.GtkCornerType(gtk_corner_window_placement))
}

func (self *ScrolledWindow) UnsetPlacement() {
	C.gtk_scrolled_window_unset_placement(self.object)
}

func (self *ScrolledWindow) SetShadowType(gtk_shadow int) {
	C.gtk_scrolled_window_set_shadow_type(self.object, C.GtkShadowType(gtk_shadow))
}

func (self *ScrolledWindow) SetHadjustment(ha *Adjustment) {
	C.gtk_scrolled_window_set_hadjustment(self.object, ha.object)
}

func (self *ScrolledWindow) SetVadjustment(va *Adjustment) {
	C.gtk_scrolled_window_set_vadjustment(self.object, va.object)
}

func (self *ScrolledWindow) GetPlacement() int {
	return int(C.gtk_scrolled_window_get_placement(self.object))
}

func (self *ScrolledWindow) GetPolicy() (hscrollbarPolicy, vscrollbarPolicy int) {
	var hp C.GtkPolicyType
	var vp C.GtkPolicyType
	C.gtk_scrolled_window_get_policy(self.object, &hp, &vp)
	return int(hp), int(vp)
}

func (self *ScrolledWindow) GetShadowType() int {
	return int(C.gtk_scrolled_window_get_shadow_type(self.object))
}

func (self *ScrolledWindow) GetMinContentWidth() int {
	return int(C.gtk_scrolled_window_get_min_content_width(self.object))
}

func (self *ScrolledWindow) SetMinContentWidth(width int) {
	C.gtk_scrolled_window_set_min_content_width(self.object, C.gint(width))
}

func (self *ScrolledWindow) GetMinContentHeight() int {
	return int(C.gtk_scrolled_window_get_min_content_height(self.object))
}

func (self *ScrolledWindow) SetMinContentHeight(height int) {
	C.gtk_scrolled_window_set_min_content_height(self.object, C.gint(height))
}

//////////////////////////////
// END GtkScrolledWindow
////////////////////////////// }}}

// End Scrolling }}}

// Miscellaneous {{{

// GtkAdjustment {{{
//////////////////////////////

// GtkAdjustment type
type Adjustment struct {
	object *C.GtkAdjustment
}

func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *Adjustment {
	a := &Adjustment{}
	o := C.gtk_adjustment_new(C.gdouble(value), C.gdouble(lower), C.gdouble(upper), C.gdouble(stepIncrement),
		C.gdouble(pageIncrement), C.gdouble(pageSize))
	a.object = C.to_GtkAdjustment(unsafe.Pointer(o))

	if gobject.IsObjectFloating(a) {
		gobject.RefSink(a)
	}
	adjustmentFinalizer(a)

	return a
}

// Clear Adjustment struct when it goes out of reach
func adjustmentFinalizer(a *Adjustment) {
	runtime.SetFinalizer(a, func(a *Adjustment) { gobject.Unref(a) })
}

// Conversion functions for gobject registration map
func newAdjustmentFromNative(obj unsafe.Pointer) interface{} {
	a := &Adjustment{}
	a.object = C.to_GtkAdjustment(obj)

	if gobject.IsObjectFloating(a) {
		gobject.RefSink(a)
	} else {
		gobject.Ref(a)
	}
	adjustmentFinalizer(a)

	return a
}

func nativeFromAdjustment(a interface{}) *gobject.GValue {
	adj, ok := a.(*Adjustment)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ADJUSTMENT, adj.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self Adjustment) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Adjustment) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Adjustment) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Adjustment) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Adjustment interface
func (self *Adjustment) GetValue() float64 {
	return float64(C.gtk_adjustment_get_value(self.object))
}

func (self *Adjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(self.object, C.gdouble(value))
}

func (self *Adjustment) ClampPage(lower, upper float64) {
	C.gtk_adjustment_clamp_page(self.object, C.gdouble(lower), C.gdouble(upper))
}

func (self *Adjustment) Changed() {
	C.gtk_adjustment_changed(self.object)
}

func (self *Adjustment) ValueChanged() {
	C.gtk_adjustment_value_changed(self.object)
}

func (self *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	C.gtk_adjustment_configure(self.object, C.gdouble(value), C.gdouble(lower), C.gdouble(upper),
		C.gdouble(stepIncrement), C.gdouble(pageIncrement), C.gdouble(pageSize))
}

func (self *Adjustment) GetLower() float64 {
	return float64(C.gtk_adjustment_get_lower(self.object))
}

func (self *Adjustment) GetPageIncrement() float64 {
	return float64(C.gtk_adjustment_get_page_increment(self.object))
}

func (self *Adjustment) GetPageSize() float64 {
	return float64(C.gtk_adjustment_get_page_size(self.object))
}

func (self *Adjustment) GetStepIncrement() float64 {
	return float64(C.gtk_adjustment_get_step_increment(self.object))
}

func (self *Adjustment) GetMinimumIncrement() float64 {
	return float64(C._gtk_adjustment_get_minimum_increment(self.object))
}

func (self *Adjustment) GetUpper() float64 {
	return float64(C.gtk_adjustment_get_upper(self.object))
}

func (self *Adjustment) SetLower(lower float64) {
	C.gtk_adjustment_set_lower(self.object, C.gdouble(lower))
}

func (self *Adjustment) SetPageIncrement(pageIncrement float64) {
	C.gtk_adjustment_set_page_increment(self.object, C.gdouble(pageIncrement))
}

func (self *Adjustment) SetPageSize(pageSize float64) {
	C.gtk_adjustment_set_page_size(self.object, C.gdouble(pageSize))
}

func (self *Adjustment) SetStepIncrement(stepIncrement float64) {
	C.gtk_adjustment_set_step_increment(self.object, C.gdouble(stepIncrement))
}

func (self *Adjustment) SetUpper(upper float64) {
	C.gtk_adjustment_set_upper(self.object, C.gdouble(upper))
}

//////////////////////////////
// GtkAdjustment
////////////////////////////// }}}

// Closure add remove funcs to/from _closures map {{{
///////////////////////////////////////////////////////////

func addToClosures(key int64, f gobject.ClosureFunc) {
	_closures[key] = f
}

func removeFromClosures(key int64) {
	if _, ok := _closures[key]; ok {
		delete(_closures, key)
	}
}

//////////////////////////////
// End Closure add/remove
////////////////////////////// }}}

// End Miscellaneous }}}

// GtkApplication {{{
//////////////////////////////

// Application Type
type Application struct {
	object *C.GtkApplication
}

// Create new application
func NewApplication(id string, flags int) *Application {
	cid := gobject.GString(id)
	defer cid.Free()
	app := C.gtk_application_new((*C.gchar)(cid.GetPtr()), C.GApplicationFlags(flags))
	gtkapp := &Application{app}

	if gobject.IsObjectFloating(gtkapp) {
		gobject.RefSink(gtkapp)
	}
	C.g_application_register(C.to_GApplication(app), nil, nil)
	//C.g_application_activate(C.to_GApplication(app))
	gtkapp.Connect("activate", func() {})

	return gtkapp
}

// Release reference when this Application struct goes out of scope
func appFinalizer(a *Application) {
	runtime.SetFinalizer(a, func(a *Application) { gobject.Unref(a) })
}

// Convert from Native to Go type
func appFromNative(app unsafe.Pointer) interface{} {
	ga := C.to_GtkApplication(app)
	gtkapp := &Application{ga}

	if gobject.IsObjectFloating(gtkapp) {
		gobject.RefSink(gtkapp)
	} else {
		gobject.Ref(gtkapp)
	}
	return &Application{ga}
}

// Convert from Go to Native
func nativeFromApp(app interface{}) *gobject.GValue {
	argapp, ok := app.(*Application)
	if ok {
		gv := gobject.CreateCGValue(GtkType.APPLICATION, argapp.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self Application) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Application) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Application) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Application) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Add Window to application object
func (self *Application) AddWindow(window *Window) {
	if window == nil {
		return
	}
	C.gtk_application_add_window(self.object, window.object)
}

// Remove Window
func (self *Application) RemoveWindow(window *Window) {
	if window == nil {
		return
	}

	C.gtk_application_remove_window(self.object, window.object)
}

// Run app
func (self *Application) Run() {
	C.run_app(self.object)
}

//////////////////////////////
// END GtkApplication
////////////////////////////// }}}

// Exported functions {{{
//////////////////////////////

//export _gtk_callback
func _gtk_callback(widget *C.GtkWidget, data C.gpointer) {
	if w, err := gobject.ConvertToGo(unsafe.Pointer(widget)); err == nil {
		id := int64(*((*C.gint64)(data)))
		_closures[id]([]interface{}{w})
	}

}

//export _gtk_entry_completion_match_func
func _gtk_entry_completion_match_func(completion *C.GtkEntryCompletion, key C.constgchar,
                iter *C.GtkTreeIter, userData C.gpointer) C.gboolean {
	entryCompletion, _ := gobject.ConvertToGo(unsafe.Pointer(completion))
	s := gobject.GoString(unsafe.Pointer(key))
	inIter := &TreeIter{*iter}

	closureId := *((*C.gint64)(userData))

	var res bool
	if c, ok := _closures[int64(closureId)]; ok {
		res = c([]interface{}{entryCompletion, s, inIter})
	}
	b := gobject.GBool(res)
	defer b.Free()

	return *((*C.gboolean)(b.GetPtr()))
}

//export _gtk_text_tag_table_foreach_func
func _gtk_text_tag_table_foreach_func(tag *C.GtkTextTag, data C.gpointer) {
	if t, err := gobject.ConvertToGo(unsafe.Pointer(tag)); err == nil {
		id := int64(*((*C.gint64)(data)))
		_closures[id]([]interface{}{t})
	}

}

//export _gtk_text_char_predicate
func _gtk_text_char_predicate(ch C.gunichar, data C.gpointer) C.gboolean {
	id := int64(*((*C.gint64)(data)))

	var res bool
	if f, ok := _closures[id]; ok {
		res = f([]interface{}{rune(ch)})
	}
	b := gobject.GBool(res)
	defer b.Free()

	return *((*C.gboolean)(b.GetPtr()))
}

//export _gtk_cell_callback
func _gtk_cell_callback(renderer *C.GtkCellRenderer, data C.gpointer) C.gboolean {
	id := int64(*((*C.gint64)(data)))

	rend, _ := gobject.ConvertToGo(unsafe.Pointer(renderer))

	var res bool
	if c, ok := _closures[id]; ok {
		res = c([]interface{}{rend})
	}
	b := gobject.GBool(res)
	defer b.Free()

	return *((*C.gboolean)(b.GetPtr()))
}

//export _gtk_cell_layout_data_func
func _gtk_cell_layout_data_func(cellLayout C.goCellLayout, cell *C.GtkCellRenderer, treeModel C.goTreeModel,
                    iter *C.GtkTreeIter, data C.gpointer) {
	id := *((*C.gint64)(data))

	goCellLayout, _ := gobject.ConvertToGo(unsafe.Pointer(cellLayout))
	goCell, _ := gobject.ConvertToGo(unsafe.Pointer(cell))
	goTreeModel, _ := gobject.ConvertToGo(unsafe.Pointer(treeModel))
	inIter := &TreeIter{*iter}

	if c, ok := _closures[int64(id)]; ok {
		c([]interface{}{goCellLayout, goCell, goTreeModel, inIter})
	}
}

//export _gtk_tree_view_row_separator_func
func _gtk_tree_view_row_separator_func(model C.goTreeModel, iter *C.GtkTreeIter, data C.gpointer) C.gboolean {
	id := *((*C.gint64)(data))
	m, _ := gobject.ConvertToGo(unsafe.Pointer(model))
	inIter := &TreeIter{*iter}

	var res bool
	if c, ok := _closures[int64(id)]; ok {
		res = c([]interface{}{m, inIter})
	}
	b := gobject.GBool(res)
	defer b.Free()

	return *((*C.gboolean)(b.GetPtr()))
}

//export _g_gtk_destroy_notify
func _g_gtk_destroy_notify(data C.gpointer) {
	id := *((*C.gint64)(data))
	if _, ok := _closures[int64(id)]; ok {
		delete(_closures, int64(id))
	}
	C.free(unsafe.Pointer(data))
}

//export _gtk_marshal
func _gtk_marshal(closure *C.GClosure,
	returnValue *C.GValue,
	n_param_values C.guint,
	paramValues C.constgvalue,
	invocationHint C.gpointer,
	marshalData C.gpointer) {

	id := int64(*((*C.gint64)(closure.data)))

	argslice := make([]interface{}, int(n_param_values))
	array := C.g_array_from_GValues(unsafe.Pointer(paramValues), n_param_values)
	for i := 0; i < int(n_param_values); i++ {
		v := C.get_index(array, C.guint(i))
		gv := gobject.NewGValueFromNative(unsafe.Pointer(&v))
		gv.ReInitializeType()
		t, e := gobject.ConvertToGo(gv.GetPtr(), gv.GetTypeID())
		if e == nil {
			argslice[i] = t
		} else {
			argslice[i] = nil
		}
	}

	if h, ok := _closures[id]; ok {
		h(argslice)
	}

	C.free_array(array)
}

//export _gtk_closure_destroy_id
func _gtk_closure_destroy_id(data C.gpointer, closure *C.GClosure) {
	id := int64(*((*C.gint64)(data)))
	if _, ok := _closures[id]; ok {
		delete(_closures, id)
	}

	C.free(unsafe.Pointer(data))
}

//export _gtk_menu_position_func
func _gtk_menu_position_func(menu *C.GtkMenu, x, y *C.gint, pushIn *C.gboolean, userData C.gpointer) {
	id := int64(*((*C.gint64)(userData)))

	goMenu := newMenuFromNative(unsafe.Pointer(menu)).(*Menu)

	var retX, retY int
	var retPush bool

	if c, ok := _closures[id]; ok {
		c([]interface{}{goMenu, &retX, &retY, &retPush})

		*x = C.gint(retX)
		*y = C.gint(retY)

		b := gobject.GBool(retPush)
		defer b.Free()
		*pushIn = *((*C.gboolean)(b.GetPtr()))
	}
}

//export _gtk_menu_detach_func
func _gtk_menu_detach_func(attachWidget *C.GtkWidget, menu *C.GtkMenu) {
	// Get id
	s := gobject.GString("detachFunc")
	defer s.Free()
	o := C.to_GObject(unsafe.Pointer(menu))
	pId := C.g_object_get_data(o, (*C.gchar)(s.GetPtr()))
	id := int64(*((*C.gint64)(pId)))

	w, _ := gobject.ConvertToGo(unsafe.Pointer(attachWidget))
	m, _ := gobject.ConvertToGo(unsafe.Pointer(menu))

	if c, ok := _closures[id]; ok {
		c([]interface{}{w, m})
	}
}

//export _gtk_tree_selection_func
func _gtk_tree_selection_func(selection *C.GtkTreeSelection, model C.goTreeModel, path C.goTreePath,
	pathCurrSelected C.gboolean, data C.gpointer) C.gboolean {

	id := int64(*((*C.gint64)(data)))

	goSelection := newTreeSelectionFromNative(unsafe.Pointer(selection)).(*TreeSelection)
	goModel := newTreeModelFromNative(unsafe.Pointer(model)).(TreeModelLike)
	goPath := newTreePathFromNative(unsafe.Pointer(path)).(*TreePath)
	goCurrSel := gobject.GoBool(unsafe.Pointer(&pathCurrSelected))

	var res bool = true
	if c, ok := _closures[id]; ok {
		res = c([]interface{}{goSelection, goModel, goPath, goCurrSel})
	}

	cret := gobject.GBool(res)
	defer cret.Free()

	return *((*C.gboolean)(cret.GetPtr()))
}

//////////////////////////////
// END Exported functions
////////////////////////////// }}}

// GtkTypes {{{

// GtkType {{{
var GtkType gtkTypes

type gtkTypes struct {
	ABOUT_DIALOG              g.GType
	ACCEL_GROUP               g.GType
	ACCEL_LABEL               g.GType
	ACCEL_MAP                 g.GType
	ACCESSIBLE                g.GType
	ACTION_GROUP              g.GType
	ACTION                    g.GType
	ACTIVATABLE               g.GType
	ADJUSTMENT                g.GType
	ALIGNMENT                 g.GType
	APP_CHOOSER_BUTTON        g.GType
	APP_CHOOSER_DIALOG        g.GType
	APP_CHOOSER               g.GType
	APP_CHOOSER_WIDGET        g.GType
	APPLICATION               g.GType
	ARROW                     g.GType
	ASPECT_FRAME              g.GType
	ASSISTANT                 g.GType
	BUTTON_BOX                g.GType
	BIN                       g.GType
	BORDER                    g.GType
	BOX                       g.GType
	BUILDABLE                 g.GType
	BUILDER                   g.GType
	BUTTON                    g.GType
	CALENDAR                  g.GType
	CELL_AREA_BOX             g.GType
	CELL_AREA_CONTEXT         g.GType
	CELL_AREA                 g.GType
	CELL_EDITABLE             g.GType
	CELL_LAYOUT               g.GType
	CELL_RENDERER_ACCEL       g.GType
	CELL_RENDERER_COMBO       g.GType
	CELL_RENDERER             g.GType
	CELL_RENDERER_PIXBUF      g.GType
	CELL_RENDERER_PROGRESS    g.GType
	CELL_RENDERER_SPIN        g.GType
	CELL_RENDERER_SPINNER     g.GType
	CELL_RENDERER_TEXT        g.GType
	CELL_RENDERER_TOGGLE      g.GType
	CELL_VIEW                 g.GType
	CHECK_BUTTON              g.GType
	CHECK_MENU_ITEM           g.GType
	CLIPBOARD                 g.GType
	COLOR_BUTTON              g.GType
	COLOR_SELECTION_DIALOG    g.GType
	COLOR_SELECTION           g.GType
	COMBO_BOX                 g.GType
	COMBO_BOX_TEXT            g.GType
	CONTAINER                 g.GType
	CSS_PROVIDER              g.GType
	DIALOG                    g.GType
	DRAWING_AREA              g.GType
	EDITABLE                  g.GType
	ENTRY_BUFFER              g.GType
	ENTRY_COMPLETION          g.GType
	ENTRY                     g.GType
	EVENT_BOX                 g.GType
	EXPANDER                  g.GType
	FILE_CHOOSER_BUTTON       g.GType
	FILE_CHOOSER_DIALOG       g.GType
	FILE_CHOOSER              g.GType
	FILE_CHOOSER_WIDGET       g.GType
	FILE_FILTER               g.GType
	FIXED                     g.GType
	FONT_BUTTON               g.GType
	FONT_SELECTION            g.GType
	FONT_SELECTION_DIALOG     g.GType
	FRAME                     g.GType
	GRADIENT                  g.GType
	GRID                      g.GType
	HANDLE_BOX                g.GType
	HBUTTON_BOX               g.GType
	HBOX                      g.GType
	HPANED                    g.GType
	HSCALE                    g.GType
	HSCROLLBAR                g.GType
	HSEPARATOR                g.GType
	HSV                       g.GType
	ICON_FACTORY              g.GType
	ICON_SET                  g.GType
	ICON_SOURCE               g.GType
	ICON_INFO                 g.GType
	ICON_THEME                g.GType
	ICON_VIEW                 g.GType
	IMAGE                     g.GType
	IMAGE_MENU_ITEM           g.GType
	IM_CONTEXT                g.GType
	IM_CONTEXT_SIMPLE         g.GType
	IM_MULTICONTEXT           g.GType
	INFO_BAR                  g.GType
	INVISIBLE                 g.GType
	LABEL                     g.GType
	LAYOUT                    g.GType
	LINK_BUTTON               g.GType
	LIST_STORE                g.GType
	MENU_BAR                  g.GType
	MENU                      g.GType
	MENU_ITEM                 g.GType
	MENU_SHELL                g.GType
	MENU_TOOL_BUTTON          g.GType
	MESSAGE_DIALOG            g.GType
	MISC                      g.GType
	MOUNT_OPERATION           g.GType
	NOTEBOOK                  g.GType
	NUMERABLE_ICON            g.GType
	OFFSCREEN_WINDOW          g.GType
	ORIENTABLE                g.GType
	PAGE_SETUP                g.GType
	PANED                     g.GType
	PAPER_SIZE                g.GType
	PLUG                      g.GType
	PRINT_CONTEXT             g.GType
	PRINT_OPERATION           g.GType
	PRINT_OPERATION_PREVIEW   g.GType
	PRINT_SETTINGS            g.GType
	PROGRESS_BAR              g.GType
	RADIO_ACTION              g.GType
	RADIO_BUTTON              g.GType
	RADIO_MENU_ITEM           g.GType
	RADIO_TOOL_BUTTON         g.GType
	RANGE                     g.GType
	RC_STYLE                  g.GType
	RECENT_ACTION             g.GType
	RECENT_CHOOSER_DIALOG     g.GType
	RECENT_CHOOSER            g.GType
	RECENT_CHOOSER_MENU       g.GType
	RECENT_CHOOSER_WIDGET     g.GType
	RECENT_FILTER             g.GType
	RECENT_INFO               g.GType
	RECENT_MANAGER            g.GType
	SCALE_BUTTON              g.GType
	SCALE                     g.GType
	SCROLLABLE                g.GType
	SCROLLBAR                 g.GType
	SCROLLED_WINDOW           g.GType
	SELECTION_DATA            g.GType
	TARGET_LIST               g.GType
	SEPARATOR                 g.GType
	SEPARATOR_MENU_ITEM       g.GType
	SEPARATOR_TOOL_ITEM       g.GType
	SETTINGS                  g.GType
	SIZE_GROUP                g.GType
	SOCKET                    g.GType
	SPIN_BUTTON               g.GType
	SPINNER                   g.GType
	STATUSBAR                 g.GType
	STATUS_ICON               g.GType
	STYLE_CONTEXT             g.GType
	STYLE                     g.GType
	STYLE_PROPERTIES          g.GType
	STYLE_PROVIDER            g.GType
	SWITCH                    g.GType
	SYMBOLIC_COLOR            g.GType
	TABLE                     g.GType
	TEAROFF_MENU_ITEM         g.GType
	TEXT_ATTRIBUTES           g.GType
	TEXT_BUFFER               g.GType
	TEXT_CHILD_ANCHOR         g.GType
	TEXT_ITER                 g.GType
	TEXT_MARK                 g.GType
	TEXT_TAG                  g.GType
	TEXT_TAG_TABLE            g.GType
	TEXT_VIEW                 g.GType
	THEMING_ENGINE            g.GType
	TOGGLE_ACTION             g.GType
	TOGGLE_BUTTON             g.GType
	TOGGLE_TOOL_BUTTON        g.GType
	TOOLBAR                   g.GType
	TOOL_BUTTON               g.GType
	TOOL_ITEM_GROUP           g.GType
	TOOL_ITEM                 g.GType
	TOOL_PALETTE              g.GType
	TOOL_SHELL                g.GType
	TOOLTIP                   g.GType
	TREE_DRAG_SOURCE          g.GType
	TREE_DRAG_DEST            g.GType
	TREE_MODEL_FILTER         g.GType
	TREE_MODEL                g.GType
	TREE_ITER                 g.GType
	TREE_PATH                 g.GType
	TREE_ROW_REFERENCE        g.GType
	TREE_MODEL_SORT           g.GType
	TREE_SELECTION            g.GType
	TREE_SORTABLE             g.GType
	TREE_STORE                g.GType
	TREE_VIEW_COLUMN          g.GType
	TREE_VIEW                 g.GType
	UI_MANAGER                g.GType
	VBUTTON_BOX               g.GType
	VBOX                      g.GType
	VIEWPORT                  g.GType
	VOLUME_BUTTON             g.GType
	VPANED                    g.GType
	VSCALE                    g.GType
	VSCROLLBAR                g.GType
	VSEPARATOR                g.GType
	WIDGET                    g.GType
	REQUISITION               g.GType
	WIDGET_PATH               g.GType
	WINDOW                    g.GType
	WINDOW_GROUP              g.GType
	LICENSE                   g.GType
	ACCEL_FLAGS               g.GType
	ASSISTANT_PAGE_TYPE       g.GType
	BUILDER_ERROR             g.GType
	CALENDAR_DISPLAY_OPTIONS  g.GType
	CELL_RENDERER_STATE       g.GType
	CELL_RENDERER_MODE        g.GType
	CELL_RENDERER_ACCEL_MODE  g.GType
	CSS_PROVIDER_ERROR        g.GType
	DEBUG_FLAG                g.GType
	DIALOG_FLAGS              g.GType
	RESPONSE_TYPE             g.GType
	DEST_DEFAULTS             g.GType
	TARGET_FLAGS              g.GType
	ENTRY_ICON_POSITION       g.GType
	ALIGN                     g.GType
	ARROW_PLACEMENT           g.GType
	ARROW_TYPE                g.GType
	ATTACH_OPTIONS            g.GType
	BUTTON_BOX_STYLE          g.GType
	DELETE_TYPE               g.GType
	DIRECTION_TYPE            g.GType
	EXPANDER_STYLE            g.GType
	ICON_SIZE                 g.GType
	SENSITIVITY_TYPE          g.GType
	TEXT_DIRECTION            g.GType
	JUSTIFICATION             g.GType
	MENU_DIRECTION_TYPE       g.GType
	MESSAGE_TYPE              g.GType
	MOVEMENT_STEP             g.GType
	SCROLL_STEP               g.GType
	ORIENTATION               g.GType
	CORNER_TYPE               g.GType
	PACK_TYPE                 g.GType
	PATH_PRIORITY_TYPE        g.GType
	PATH_TYPE                 g.GType
	POLICY_TYPE               g.GType
	POSITION_TYPE             g.GType
	RELIEF_STYLE              g.GType
	RESIZE_MODE               g.GType
	SCROLL_TYPE               g.GType
	SELECTION_MODE            g.GType
	SHADOW_TYPE               g.GType
	STATE_TYPE                g.GType
	TOOLBAR_STYLE             g.GType
	WINDOW_POSITION           g.GType
	WINDOW_TYPE               g.GType
	WRAP_MODE                 g.GType
	SORT_TYPE                 g.GType
	IM_PREEDIT_STYLE          g.GType
	IM_STATUS_STYLE           g.GType
	PACK_DIRECTION            g.GType
	PRINT_PAGES               g.GType
	PAGE_SET                  g.GType
	NUMBER_UP_LAYOUT          g.GType
	PAGE_ORIENTATION          g.GType
	PRINT_QUALITY             g.GType
	PRINT_DUPLEX              g.GType
	UNIT                      g.GType
	TREE_VIEW_GRID_LINES      g.GType
	DRAG_RESULT               g.GType
	SIZE_REQUEST_MODE         g.GType
	SCROLLABLE_POLICY         g.GType
	STATE_FLAGS               g.GType
	REGION_FLAGS              g.GType
	JUNCTION_SIDES            g.GType
	BORDER_STYLE              g.GType
	FILE_CHOOSER_ACTION       g.GType
	FILE_CHOOSER_CONFIRMATION g.GType
	FILE_CHOOSER_ERROR        g.GType
	FILE_FILTER_FLAGS         g.GType
	ICON_LOOKUP_FLAGS         g.GType
	ICON_THEME_ERROR          g.GType
	ICON_VIEW_DROP_POSITION   g.GType
	IMAGE_TYPE                g.GType
	BUTTONS_TYPE              g.GType
	NOTEBOOK_TAB              g.GType
	PRINT_STATUS              g.GType
	PRINT_OPERATION_RESULT    g.GType
	PRINT_OPERATION_ACTION    g.GType
	PRINT_ERROR               g.GType
	RC_FLAGS                  g.GType
	RC_TOKEN_TYPE             g.GType
	RECENT_SORT_TYPE          g.GType
	RECENT_CHOOSER_ERROR      g.GType
	RECENT_FILTER_FLAGS       g.GType
	RECENT_MANAGER_ERROR      g.GType
	SIZE_GROUP_MODE           g.GType
	SPIN_BUTTON_UPDATE_POLICY g.GType
	SPIN_TYPE                 g.GType
	TEXT_BUFFER_TARGET_INFO   g.GType
	TEXT_SEARCH_FLAGS         g.GType
	TEXT_WINDOW_TYPE          g.GType
	TOOLBAR_SPACE_STYLE       g.GType
	TOOL_PALETTE_DRAG_TARGETS g.GType
	TREE_MODEL_FLAGS          g.GType
	TREE_VIEW_DROP_POSITION   g.GType
	TREE_VIEW_COLUMN_SIZING   g.GType
	UI_MANAGER_ITEM_TYPE      g.GType
	WIDGET_HELP_TYPE          g.GType
} // }}}

// GtkShadowType {{{
var GtkShadow gtkShadow

type gtkShadow struct {
	NONE       int
	IN         int
	OUT        int
	ETCHED_IN  int
	ETCHED_OUT int
}

// End GtkShadow }}}

// GtkStock {{{
var GtkStock gtkStock

type gtkStock struct {
	ABOUT                         string
	ADD                           string
	APPLY                         string
	BOLD                          string
	CANCEL                        string
	CAPS_LOCK_WARNING             string
	CDROM                         string
	CLEAR                         string
	CLOSE                         string
	COLOR_PICKER                  string
	CONNECT                       string
	CONVERT                       string
	COPY                          string
	CUT                           string
	DELETE                        string
	DIALOG_AUTHENTICATION         string
	DIALOG_INFO                   string
	DIALOG_WARNING                string
	DIALOG_ERROR                  string
	DIALOG_QUESTION               string
	DIRECTORY                     string
	DISCARD                       string
	DISCONNECT                    string
	DND                           string
	DND_MULTIPLE                  string
	EDIT                          string
	EXECUTE                       string
	FILE                          string
	FIND                          string
	FIND_AND_REPLACE              string
	FLOPPY                        string
	FULLSCREEN                    string
	GOTO_BOTTOM                   string
	GOTO_FIRST                    string
	GOTO_LAST                     string
	GOTO_TOP                      string
	GO_BACK                       string
	GO_DOWN                       string
	GO_FORWARD                    string
	GO_UP                         string
	HARDDISK                      string
	HELP                          string
	HOME                          string
	INDEX                         string
	INDENT                        string
	INFO                          string
	ITALIC                        string
	JUMP_TO                       string
	JUSTIFY_CENTER                string
	JUSTIFY_FILL                  string
	JUSTIFY_LEFT                  string
	JUSTIFY_RIGHT                 string
	LEAVE_FULLSCREEN              string
	MISSING_IMAGE                 string
	MEDIA_FORWARD                 string
	MEDIA_NEXT                    string
	MEDIA_PAUSE                   string
	MEDIA_PLAY                    string
	MEDIA_PREVIOUS                string
	MEDIA_RECORD                  string
	MEDIA_REWIND                  string
	MEDIA_STOP                    string
	NETWORK                       string
	NEW                           string
	NO                            string
	OK                            string
	OPEN                          string
	ORIENTATION_PORTRAIT          string
	ORIENTATION_LANDSCAPE         string
	ORIENTATION_REVERSE_LANDSCAPE string
	ORIENTATION_REVERSE_PORTRAIT  string
	PAGE_SETUP                    string
	PASTE                         string
	PREFERENCES                   string
	PRINT                         string
	PRINT_ERROR                   string
	PRINT_PAUSED                  string
	PRINT_PREVIEW                 string
	PRINT_REPORT                  string
	PRINT_WARNING                 string
	PROPERTIES                    string
	QUIT                          string
	REDO                          string
	REFRESH                       string
	REMOVE                        string
	REVERT_TO_SAVED               string
	SAVE                          string
	SAVE_AS                       string
	SELECT_ALL                    string
	SELECT_COLOR                  string
	SELECT_FONT                   string
	SORT_ASCENDING                string
	SORT_DESCENDING               string
	SPELL_CHECK                   string
	STOP                          string
	STRIKETHROUGH                 string
	UNDELETE                      string
	UNDERLINE                     string
	UNDO                          string
	UNINDENT                      string
	YES                           string
	ZOOM_100                      string
	ZOOM_FIT                      string
	ZOOM_IN                       string
	ZOOM_OUT                      string
}

// End GtkStock  }}}

// GtkPositionType  {{{
var GtkPosition gtkPosition

type gtkPosition struct {
	LEFT   int
	RIGHT  int
	TOP    int
	BOTTOM int
}

// End GtkPositionType }}}

// GtkDialogFlags {{{
var GtkDialogFlags gtkDialogFlags

type gtkDialogFlags struct {
	MODAL               int
	DESTROY_WITH_PARENT int
}

// End GtkDialogFlags }}}

// GtkResponse {{{
var GtkResponse gtkResponse

type gtkResponse struct {
	NONE         int
	REJECT       int
	ACCEPT       int
	DELETE_EVENT int
	OK           int
	CANCEL       int
	CLOSE        int
	YES          int
	NO           int
	APPLY        int
	HELP         int
}

// End GtkResponse }}}

// GtkMessageType {{{
var GtkMessage gtkMessage

type gtkMessage struct {
	INFO     int
	WARNING  int
	QUESTION int
	ERROR    int
	OTHER    int
}

// End GtkMessageType }}}

// GtkButtons {{{
var GtkButtons gtkButtons

type gtkButtons struct {
	NONE      int
	OK        int
	CLOSE     int
	CANCEL    int
	YES_NO    int
	OK_CANCEL int
}

// End GtkButtons }}}

// GtkOrientation {{{
var GtkOrientation gtkOrientation

type gtkOrientation struct {
	HORIZONTAL int
	VERTICAL   int
}

// End GtkOrientation }}}

// GtkPackType {{{
var GtkPack gtkPack

type gtkPack struct {
	PACK_START int
	PACK_END   int
}

// End GtkPackType }}}

// GtkButtonBoxStyle {{{
var GtkButtonBoxStyle gtkButtonBoxStyle

type gtkButtonBoxStyle struct {
	SPREAD int
	EDGE   int
	START  int
	END    int
	CENTER int
}

// End GtkButtonBoxStyle }}}

// GtkIconSize {{{
var GtkIconSize gtkIconSize

type gtkIconSize struct {
	INVALID       int
	MENU          int
	SMALL_TOOLBAR int
	LARGE_TOOLBAR int
	BUTTON        int
	DND           int
	DIALOG        int
}

// End GtkIconSize  }}}

// GApplicatonFlags {{{
var GApplicationFlags gApplicationFlags

type gApplicationFlags struct {
	FLAGS_NONE           int
	IS_SERVICE           int
	IS_LAUNCHER          int
	HANDLES_OPEN         int
	HANDLES_COMMAND_LINE int
	SEND_ENVIRONMENT     int
	NON_UNIQUE           int
}

// End GApplicationFlags }}}

// GtkJustification {{{
var GtkJustification gtkJustification

type gtkJustification struct {
	LEFT   int
	RIGHT  int
	CENTER int
	FILL   int
}

// End GtkJustification }}}

// GtkWindowType {{{
var GtkWindowType gtkWindowType

type gtkWindowType struct {
	TOPLEVEL int
	POPUP    int
}

// End GtkWindowType }}}

// GtkSensitivityType {{{
var GtkSensitivity gtkSensitivity

type gtkSensitivity struct {
	AUTO int
	ON   int
	OFF  int
}

// End GtkSensitivity }}}

// GtkPolicyType {{{
var GtkPolicy gtkPolicy

type gtkPolicy struct {
	ALWAYS    int
	AUTOMATIC int
	NEVER     int
}

// End GtkPolicyType }}}

// GtkCornerType {{{
var GtkCorner gtkCorner

type gtkCorner struct {
	TOP_LEFT     int
	BOTTOM_LEFT  int
	TOP_RIGHT    int
	BOTTOM_RIGHT int
}

// End GtkCornerType }}}

// GtkTreeModelFlags  {{{
var GtkTreeModelFlags gtkTreeModelFlags

type gtkTreeModelFlags struct {
	ITERS_PERSIST int
	LIST_ONLY     int
}

// End GtkTreeModelFlags  }}}

// GtkCellRendererState {{{
var GtkCellRendererState gtkCellRendererState

type gtkCellRendererState struct {
	SELECTED    int
	PRELIT      int
	INSENSITIVE int
	SORTED      int
	FOCUSED     int
}

// End GtkCellRendererState }}}

// GtkCellRendererMode {{{
var GtkCellRendererMode gtkCellRendererMode

type gtkCellRendererMode struct {
	INERT       int
	ACTIVATABLE int
	EDITABLE    int
}

// End GtkCellRendererMode }}}

// GtkSizeRequestMode {{{
var GtkSizeRequestMode gtkSizeRequestMode

type gtkSizeRequestMode struct {
	HEIGHT_FOR_WIDTH int
	WIDTH_FOR_HEIGHT int
	CONSTANT_SIZE    int
}

// End GtkSizeRequestMode  }}}

// GtkCellRendererAccelMode {{{
var GtkCellRendererAccelMode gtkCellRendererAccelMode

type gtkCellRendererAccelMode struct {
	GTK   int
	OTHER int
}

// End GtkCellRendererAccelMode }}}

// GtkTreeViewColumnSizing {{{
var GtkTreeViewColumnSizing gtkTreeViewColumnSizing

type gtkTreeViewColumnSizing struct {
	GROW_ONLY int
	AUTOSIZE  int
	FIXED     int
}

// End GtkTreeViewColumnSizing }}}

// GtkSortType {{{
var GtkSortType gtkSortType

type gtkSortType struct {
	ASCENDING  int
	DESCENDING int
}

// End GtkSortType }}}

// GtkSelectionMode {{{
var GtkSelectionMode gtkSelectionMode

type gtkSelectionMode struct {
	NONE     int
	SINGLE   int
	BROWSE   int
	MULTIPLE int
}

// End GtkSelectionMode }}}

// GtkResizeMode {{{
var GtkResizeMode gtkResizeMode

type gtkResizeMode struct {
	PARENT int
	QUEUE  int
}

// End GtkResizeMode }}}

// GtkAssistantPage {{{
var GtkAssistantPage gtkAssistantPage

type gtkAssistantPage struct {
	CONTENT  int
	INTRO    int
	CONFIRM  int
	SUMMARY  int
	PROGRESS int
	CUSTOM   int
}

// End GtkAssistantPageType }}}

// GtkAlign {{{
var GtkAlign gtkAlign

type gtkAlign struct {
	FILL   int
	START  int
	END    int
	CENTER int
}

// End GtkAlign }}}

// GtkTextSearchFlags {{{

var GtkTextSearchFlags gtkTextSearchFlags

type gtkTextSearchFlags struct {
	VISIBLE_ONLY     int
	TEXT_ONLY        int
	CASE_INSENSITIVE int
}

// End GtkTextSearchFlags }}}

// GtkTextWindowType {{{

var GtkTextWindowType gtkTextWindowType

type gtkTextWindowType struct {
	PRIVATE int
	WIDGET  int
	TEXT    int
	LEFT    int
	RIGHT   int
	TOP     int
	BOTTOM  int
}

// End GtkTextWindowType }}}

// GtkWrapMode {{{

var GtkWrapMode gtkWrapMode

type gtkWrapMode struct {
	NONE      int
	CHAR      int
	WORD      int
	WORD_CHAR int
}

// End GtkWrapMode }}}

// GtkTextDirection {{{

var GtkTextDirection gtkTextDirection

type gtkTextDirection struct {
	NONE int
	LTR  int
	RTL  int
}

// End GtkTextDirection }}}

// GtkAccelFlags {{{

var GtkAccelFlags gtkAccelFlags

type gtkAccelFlags struct {
	VISIBLE int
	LOCKED  int
	MASK    int
}

// End GtkAccelFlags }}}

// GtkMenuDirection {{{

var GtkMenuDirection gtkMenuDirection

type gtkMenuDirection struct {
	PARENT int
	CHILD  int
	NEXT   int
	PREV   int
}

// End GtkMenuDirection }}}

// GtkPackDirection {{{

var GtkPackDirection gtkPackDirection

type gtkPackDirection struct {
	LTR int
	RTL int
	TTB int
	BTT int
}

// End GtkPackDirection }}}

// GtkSpinButtonUpdatePolicy {{{

var GtkSpinButtonUpdatePolicy gtkSpinButtonUpdatePolicy

type gtkSpinButtonUpdatePolicy struct {
	ALWAYS   int
	IF_VALID int
}

// End GtkSpinButtonUpdatePolicy }}}

// GtkSpinType {{{

var GtkSpinType gtkSpinType

type gtkSpinType struct {
	STEP_FORWARD  int
	STEP_BACKWARD int
	PAGE_FORWARD  int
	PAGE_BACKWARD int
	HOME          int
	END           int
	USER_DEFINED  int
}

// End GtkSpinType }}}

// GtkTreeViewGridLines {{{

var GtkTreeViewGridLines gtkTreeViewGridLines

type gtkTreeViewGridLines struct {
	NONE       int
	HORIZONTAL int
	VERTICAL   int
	BOTH       int
}

// End GtkTreeViewGridLines }}}

// GtkEntryIconPosition {{{

var GtkEntryIconPosition gtkEntryIconPosition

type gtkEntryIconPosition struct {
	PRIMARY   int
	SECONDARY int
}

// End GtkEntryIconPosition }}}

// GtkImageType {{{

var GtkImageType gtkImageType

type gtkImageType struct {
	EMPTY     int
	PIXBUF    int
	STOCK     int
	ICON_SET  int
	ANIMATION int
	ICON_NAME int
	GICON     int
}

// End GtkImageType }}}

// End GtkTypes }}}

// GTK3 MODULE init function {{{
func init() {
	// Initialiize map for closures
	_closures = make(map[int64]gobject.ClosureFunc)

    // Initialize GtkType struct {{{
	GtkType.ABOUT_DIALOG = g.GType(C.gtk_about_dialog_get_type())
	GtkType.ACCEL_GROUP = g.GType(C.gtk_accel_group_get_type())
	GtkType.ACCEL_LABEL = g.GType(C.gtk_accel_label_get_type())
	GtkType.ACCEL_MAP = g.GType(C.gtk_accel_map_get_type())
	GtkType.ACCESSIBLE = g.GType(C.gtk_accessible_get_type())
	GtkType.ACTION_GROUP = g.GType(C.gtk_action_group_get_type())
	GtkType.ACTION = g.GType(C.gtk_action_get_type())
	GtkType.ACTIVATABLE = g.GType(C.gtk_activatable_get_type())
	GtkType.ADJUSTMENT = g.GType(C.gtk_adjustment_get_type())
	GtkType.ALIGNMENT = g.GType(C.gtk_alignment_get_type())
	GtkType.APP_CHOOSER_BUTTON = g.GType(C.gtk_app_chooser_button_get_type())
	GtkType.APP_CHOOSER_DIALOG = g.GType(C.gtk_app_chooser_dialog_get_type())
	GtkType.APP_CHOOSER = g.GType(C.gtk_app_chooser_get_type())
	GtkType.APP_CHOOSER_WIDGET = g.GType(C.gtk_app_chooser_widget_get_type())
	GtkType.APPLICATION = g.GType(C.gtk_application_get_type())
	GtkType.ARROW = g.GType(C.gtk_arrow_get_type())
	GtkType.ASPECT_FRAME = g.GType(C.gtk_aspect_frame_get_type())
	GtkType.ASSISTANT = g.GType(C.gtk_assistant_get_type())
	GtkType.BUTTON_BOX = g.GType(C.gtk_button_box_get_type())
	GtkType.BIN = g.GType(C.gtk_bin_get_type())
	GtkType.BORDER = g.GType(C.gtk_border_get_type())
	GtkType.BOX = g.GType(C.gtk_box_get_type())
	GtkType.BUILDABLE = g.GType(C.gtk_buildable_get_type())
	GtkType.BUILDER = g.GType(C.gtk_builder_get_type())
	GtkType.BUTTON = g.GType(C.gtk_button_get_type())
	GtkType.CALENDAR = g.GType(C.gtk_calendar_get_type())
	GtkType.CELL_AREA_BOX = g.GType(C.gtk_cell_area_box_get_type())
	GtkType.CELL_AREA_CONTEXT = g.GType(C.gtk_cell_area_context_get_type())
	GtkType.CELL_AREA = g.GType(C.gtk_cell_area_get_type())
	GtkType.CELL_EDITABLE = g.GType(C.gtk_cell_editable_get_type())
	GtkType.CELL_LAYOUT = g.GType(C.gtk_cell_layout_get_type())
	GtkType.CELL_RENDERER_ACCEL = g.GType(C.gtk_cell_renderer_accel_get_type())
	GtkType.CELL_RENDERER_COMBO = g.GType(C.gtk_cell_renderer_combo_get_type())
	GtkType.CELL_RENDERER = g.GType(C.gtk_cell_renderer_get_type())
	GtkType.CELL_RENDERER_PIXBUF = g.GType(C.gtk_cell_renderer_pixbuf_get_type())
	GtkType.CELL_RENDERER_PROGRESS = g.GType(C.gtk_cell_renderer_progress_get_type())
	GtkType.CELL_RENDERER_SPIN = g.GType(C.gtk_cell_renderer_spin_get_type())
	GtkType.CELL_RENDERER_SPINNER = g.GType(C.gtk_cell_renderer_spinner_get_type())
	GtkType.CELL_RENDERER_TEXT = g.GType(C.gtk_cell_renderer_text_get_type())
	GtkType.CELL_RENDERER_TOGGLE = g.GType(C.gtk_cell_renderer_toggle_get_type())
	GtkType.CELL_VIEW = g.GType(C.gtk_cell_view_get_type())
	GtkType.CHECK_BUTTON = g.GType(C.gtk_check_button_get_type())
	GtkType.CHECK_MENU_ITEM = g.GType(C.gtk_check_menu_item_get_type())
	GtkType.CLIPBOARD = g.GType(C.gtk_clipboard_get_type())
	GtkType.COLOR_BUTTON = g.GType(C.gtk_color_button_get_type())
	GtkType.COLOR_SELECTION_DIALOG = g.GType(C.gtk_color_selection_dialog_get_type())
	GtkType.COLOR_SELECTION = g.GType(C.gtk_color_selection_get_type())
	GtkType.COMBO_BOX = g.GType(C.gtk_combo_box_get_type())
	GtkType.COMBO_BOX_TEXT = g.GType(C.gtk_combo_box_text_get_type())
	GtkType.CONTAINER = g.GType(C.gtk_container_get_type())
	GtkType.CSS_PROVIDER = g.GType(C.gtk_css_provider_get_type())
	GtkType.DIALOG = g.GType(C.gtk_dialog_get_type())
	GtkType.DRAWING_AREA = g.GType(C.gtk_drawing_area_get_type())
	GtkType.EDITABLE = g.GType(C.gtk_editable_get_type())
	GtkType.ENTRY_BUFFER = g.GType(C.gtk_entry_buffer_get_type())
	GtkType.ENTRY_COMPLETION = g.GType(C.gtk_entry_completion_get_type())
	GtkType.ENTRY = g.GType(C.gtk_entry_get_type())
	GtkType.EVENT_BOX = g.GType(C.gtk_event_box_get_type())
	GtkType.EXPANDER = g.GType(C.gtk_expander_get_type())
	GtkType.FILE_CHOOSER_BUTTON = g.GType(C.gtk_file_chooser_button_get_type())
	GtkType.FILE_CHOOSER_DIALOG = g.GType(C.gtk_file_chooser_dialog_get_type())
	GtkType.FILE_CHOOSER = g.GType(C.gtk_file_chooser_get_type())
	GtkType.FILE_CHOOSER_WIDGET = g.GType(C.gtk_file_chooser_widget_get_type())
	GtkType.FILE_FILTER = g.GType(C.gtk_file_filter_get_type())
	GtkType.FIXED = g.GType(C.gtk_fixed_get_type())
	GtkType.FONT_BUTTON = g.GType(C.gtk_font_button_get_type())
	GtkType.FONT_SELECTION = g.GType(C.gtk_font_selection_get_type())
	GtkType.FONT_SELECTION_DIALOG = g.GType(C.gtk_font_selection_dialog_get_type())
	GtkType.FRAME = g.GType(C.gtk_frame_get_type())
	GtkType.GRADIENT = g.GType(C.gtk_gradient_get_type())
	GtkType.GRID = g.GType(C.gtk_grid_get_type())
	GtkType.HANDLE_BOX = g.GType(C.gtk_handle_box_get_type())
	GtkType.HBUTTON_BOX = g.GType(C.gtk_hbutton_box_get_type())
	GtkType.HBOX = g.GType(C.gtk_hbox_get_type())
	GtkType.HPANED = g.GType(C.gtk_hpaned_get_type())
	GtkType.HSCALE = g.GType(C.gtk_hscale_get_type())
	GtkType.HSCROLLBAR = g.GType(C.gtk_hscrollbar_get_type())
	GtkType.HSEPARATOR = g.GType(C.gtk_hseparator_get_type())
	GtkType.HSV = g.GType(C.gtk_hsv_get_type())
	GtkType.ICON_FACTORY = g.GType(C.gtk_icon_factory_get_type())
	GtkType.ICON_SET = g.GType(C.gtk_icon_set_get_type())
	GtkType.ICON_SOURCE = g.GType(C.gtk_icon_source_get_type())
	GtkType.ICON_INFO = g.GType(C.gtk_icon_info_get_type())
	GtkType.ICON_THEME = g.GType(C.gtk_icon_theme_get_type())
	GtkType.ICON_VIEW = g.GType(C.gtk_icon_view_get_type())
	GtkType.IMAGE = g.GType(C.gtk_image_get_type())
	GtkType.IMAGE_MENU_ITEM = g.GType(C.gtk_image_menu_item_get_type())
	GtkType.IM_CONTEXT = g.GType(C.gtk_im_context_get_type())
	GtkType.IM_CONTEXT_SIMPLE = g.GType(C.gtk_im_context_simple_get_type())
	GtkType.IM_MULTICONTEXT = g.GType(C.gtk_im_multicontext_get_type())
	GtkType.INFO_BAR = g.GType(C.gtk_info_bar_get_type())
	GtkType.INVISIBLE = g.GType(C.gtk_invisible_get_type())
	GtkType.LABEL = g.GType(C.gtk_label_get_type())
	GtkType.LAYOUT = g.GType(C.gtk_layout_get_type())
	GtkType.LINK_BUTTON = g.GType(C.gtk_link_button_get_type())
	GtkType.LIST_STORE = g.GType(C.gtk_list_store_get_type())
	GtkType.MENU_BAR = g.GType(C.gtk_menu_bar_get_type())
	GtkType.MENU = g.GType(C.gtk_menu_get_type())
	GtkType.MENU_ITEM = g.GType(C.gtk_menu_item_get_type())
	GtkType.MENU_SHELL = g.GType(C.gtk_menu_shell_get_type())
	GtkType.MENU_TOOL_BUTTON = g.GType(C.gtk_menu_tool_button_get_type())
	GtkType.MESSAGE_DIALOG = g.GType(C.gtk_message_dialog_get_type())
	GtkType.MISC = g.GType(C.gtk_misc_get_type())
	GtkType.MOUNT_OPERATION = g.GType(C.gtk_mount_operation_get_type())
	GtkType.NOTEBOOK = g.GType(C.gtk_notebook_get_type())
	GtkType.NUMERABLE_ICON = g.GType(C.gtk_numerable_icon_get_type())
	GtkType.OFFSCREEN_WINDOW = g.GType(C.gtk_offscreen_window_get_type())
	GtkType.ORIENTABLE = g.GType(C.gtk_orientable_get_type())
	GtkType.PAGE_SETUP = g.GType(C.gtk_page_setup_get_type())
	GtkType.PANED = g.GType(C.gtk_paned_get_type())
	GtkType.PAPER_SIZE = g.GType(C.gtk_paper_size_get_type())
	GtkType.PLUG = g.GType(C.gtk_plug_get_type())
	GtkType.PRINT_CONTEXT = g.GType(C.gtk_print_context_get_type())
	GtkType.PRINT_OPERATION = g.GType(C.gtk_print_operation_get_type())
	GtkType.PRINT_OPERATION_PREVIEW = g.GType(C.gtk_print_operation_preview_get_type())
	GtkType.PRINT_SETTINGS = g.GType(C.gtk_print_settings_get_type())
	GtkType.PROGRESS_BAR = g.GType(C.gtk_progress_bar_get_type())
	GtkType.RADIO_ACTION = g.GType(C.gtk_radio_action_get_type())
	GtkType.RADIO_BUTTON = g.GType(C.gtk_radio_button_get_type())
	GtkType.RADIO_MENU_ITEM = g.GType(C.gtk_radio_menu_item_get_type())
	GtkType.RADIO_TOOL_BUTTON = g.GType(C.gtk_radio_tool_button_get_type())
	GtkType.RANGE = g.GType(C.gtk_range_get_type())
	GtkType.RC_STYLE = g.GType(C.gtk_rc_style_get_type())
	GtkType.RECENT_ACTION = g.GType(C.gtk_recent_action_get_type())
	GtkType.RECENT_CHOOSER_DIALOG = g.GType(C.gtk_recent_chooser_dialog_get_type())
	GtkType.RECENT_CHOOSER = g.GType(C.gtk_recent_chooser_get_type())
	GtkType.RECENT_CHOOSER_MENU = g.GType(C.gtk_recent_chooser_menu_get_type())
	GtkType.RECENT_CHOOSER_WIDGET = g.GType(C.gtk_recent_chooser_widget_get_type())
	GtkType.RECENT_FILTER = g.GType(C.gtk_recent_filter_get_type())
	GtkType.RECENT_INFO = g.GType(C.gtk_recent_info_get_type())
	GtkType.RECENT_MANAGER = g.GType(C.gtk_recent_manager_get_type())
	GtkType.SCALE_BUTTON = g.GType(C.gtk_scale_button_get_type())
	GtkType.SCALE = g.GType(C.gtk_scale_get_type())
	GtkType.SCROLLABLE = g.GType(C.gtk_scrollable_get_type())
	GtkType.SCROLLBAR = g.GType(C.gtk_scrollbar_get_type())
	GtkType.SCROLLED_WINDOW = g.GType(C.gtk_scrolled_window_get_type())
	GtkType.SELECTION_DATA = g.GType(C.gtk_selection_data_get_type())
	GtkType.TARGET_LIST = g.GType(C.gtk_target_list_get_type())
	GtkType.SEPARATOR = g.GType(C.gtk_separator_get_type())
	GtkType.SEPARATOR_MENU_ITEM = g.GType(C.gtk_separator_menu_item_get_type())
	GtkType.SEPARATOR_TOOL_ITEM = g.GType(C.gtk_separator_tool_item_get_type())
	GtkType.SETTINGS = g.GType(C.gtk_settings_get_type())
	GtkType.SIZE_GROUP = g.GType(C.gtk_size_group_get_type())
	GtkType.SOCKET = g.GType(C.gtk_socket_get_type())
	GtkType.SPIN_BUTTON = g.GType(C.gtk_spin_button_get_type())
	GtkType.SPINNER = g.GType(C.gtk_spinner_get_type())
	GtkType.STATUSBAR = g.GType(C.gtk_statusbar_get_type())
	GtkType.STATUS_ICON = g.GType(C.gtk_status_icon_get_type())
	GtkType.STYLE_CONTEXT = g.GType(C.gtk_style_context_get_type())
	GtkType.STYLE = g.GType(C.gtk_style_get_type())
	GtkType.STYLE_PROPERTIES = g.GType(C.gtk_style_properties_get_type())
	GtkType.STYLE_PROVIDER = g.GType(C.gtk_style_provider_get_type())
	GtkType.SWITCH = g.GType(C.gtk_switch_get_type())
	GtkType.SYMBOLIC_COLOR = g.GType(C.gtk_symbolic_color_get_type())
	GtkType.TABLE = g.GType(C.gtk_table_get_type())
	GtkType.TEAROFF_MENU_ITEM = g.GType(C.gtk_tearoff_menu_item_get_type())
	GtkType.TEXT_ATTRIBUTES = g.GType(C.gtk_text_attributes_get_type())
	GtkType.TEXT_BUFFER = g.GType(C.gtk_text_buffer_get_type())
	GtkType.TEXT_CHILD_ANCHOR = g.GType(C.gtk_text_child_anchor_get_type())
	GtkType.TEXT_ITER = g.GType(C.gtk_text_iter_get_type())
	GtkType.TEXT_MARK = g.GType(C.gtk_text_mark_get_type())
	GtkType.TEXT_TAG = g.GType(C.gtk_text_tag_get_type())
	GtkType.TEXT_TAG_TABLE = g.GType(C.gtk_text_tag_table_get_type())
	GtkType.TEXT_VIEW = g.GType(C.gtk_text_view_get_type())
	GtkType.THEMING_ENGINE = g.GType(C.gtk_theming_engine_get_type())
	GtkType.TOGGLE_ACTION = g.GType(C.gtk_toggle_action_get_type())
	GtkType.TOGGLE_BUTTON = g.GType(C.gtk_toggle_button_get_type())
	GtkType.TOGGLE_TOOL_BUTTON = g.GType(C.gtk_toggle_tool_button_get_type())
	GtkType.TOOLBAR = g.GType(C.gtk_toolbar_get_type())
	GtkType.TOOL_BUTTON = g.GType(C.gtk_tool_button_get_type())
	GtkType.TOOL_ITEM_GROUP = g.GType(C.gtk_tool_item_group_get_type())
	GtkType.TOOL_ITEM = g.GType(C.gtk_tool_item_get_type())
	GtkType.TOOL_PALETTE = g.GType(C.gtk_tool_palette_get_type())
	GtkType.TOOL_SHELL = g.GType(C.gtk_tool_shell_get_type())
	GtkType.TOOLTIP = g.GType(C.gtk_tooltip_get_type())
	GtkType.TREE_DRAG_SOURCE = g.GType(C.gtk_tree_drag_source_get_type())
	GtkType.TREE_DRAG_DEST = g.GType(C.gtk_tree_drag_dest_get_type())
	GtkType.TREE_MODEL_FILTER = g.GType(C.gtk_tree_model_filter_get_type())
	GtkType.TREE_MODEL = g.GType(C.gtk_tree_model_get_type())
	GtkType.TREE_ITER = g.GType(C.gtk_tree_iter_get_type())
	GtkType.TREE_PATH = g.GType(C.gtk_tree_path_get_type())
	GtkType.TREE_ROW_REFERENCE = g.GType(C.gtk_tree_row_reference_get_type())
	GtkType.TREE_MODEL_SORT = g.GType(C.gtk_tree_model_sort_get_type())
	GtkType.TREE_SELECTION = g.GType(C.gtk_tree_selection_get_type())
	GtkType.TREE_SORTABLE = g.GType(C.gtk_tree_sortable_get_type())
	GtkType.TREE_STORE = g.GType(C.gtk_tree_store_get_type())
	GtkType.TREE_VIEW_COLUMN = g.GType(C.gtk_tree_view_column_get_type())
	GtkType.TREE_VIEW = g.GType(C.gtk_tree_view_get_type())
	GtkType.UI_MANAGER = g.GType(C.gtk_ui_manager_get_type())
	GtkType.VBUTTON_BOX = g.GType(C.gtk_vbutton_box_get_type())
	GtkType.VBOX = g.GType(C.gtk_vbox_get_type())
	GtkType.VIEWPORT = g.GType(C.gtk_viewport_get_type())
	GtkType.VOLUME_BUTTON = g.GType(C.gtk_volume_button_get_type())
	GtkType.VPANED = g.GType(C.gtk_vpaned_get_type())
	GtkType.VSCALE = g.GType(C.gtk_vscale_get_type())
	GtkType.VSCROLLBAR = g.GType(C.gtk_vscrollbar_get_type())
	GtkType.VSEPARATOR = g.GType(C.gtk_vseparator_get_type())
	GtkType.WIDGET = g.GType(C.gtk_widget_get_type())
	GtkType.REQUISITION = g.GType(C.gtk_requisition_get_type())
	GtkType.WIDGET_PATH = g.GType(C.gtk_widget_path_get_type())
	GtkType.WINDOW = g.GType(C.gtk_window_get_type())
	GtkType.WINDOW_GROUP = g.GType(C.gtk_window_group_get_type())
	GtkType.LICENSE = g.GType(C.gtk_license_get_type())
	GtkType.ACCEL_FLAGS = g.GType(C.gtk_accel_flags_get_type())
	GtkType.ASSISTANT_PAGE_TYPE = g.GType(C.gtk_assistant_page_type_get_type())
	GtkType.BUILDER_ERROR = g.GType(C.gtk_builder_error_get_type())
	GtkType.CALENDAR_DISPLAY_OPTIONS = g.GType(C.gtk_calendar_display_options_get_type())
	GtkType.CELL_RENDERER_STATE = g.GType(C.gtk_cell_renderer_state_get_type())
	GtkType.CELL_RENDERER_MODE = g.GType(C.gtk_cell_renderer_mode_get_type())
	GtkType.CELL_RENDERER_ACCEL_MODE = g.GType(C.gtk_cell_renderer_accel_mode_get_type())
	GtkType.CSS_PROVIDER_ERROR = g.GType(C.gtk_css_provider_error_get_type())
	GtkType.DEBUG_FLAG = g.GType(C.gtk_debug_flag_get_type())
	GtkType.DIALOG_FLAGS = g.GType(C.gtk_dialog_flags_get_type())
	GtkType.RESPONSE_TYPE = g.GType(C.gtk_response_type_get_type())
	GtkType.DEST_DEFAULTS = g.GType(C.gtk_dest_defaults_get_type())
	GtkType.TARGET_FLAGS = g.GType(C.gtk_target_flags_get_type())
	GtkType.ENTRY_ICON_POSITION = g.GType(C.gtk_entry_icon_position_get_type())
	GtkType.ALIGN = g.GType(C.gtk_align_get_type())
	GtkType.ARROW_PLACEMENT = g.GType(C.gtk_arrow_placement_get_type())
	GtkType.ARROW_TYPE = g.GType(C.gtk_arrow_type_get_type())
	GtkType.ATTACH_OPTIONS = g.GType(C.gtk_attach_options_get_type())
	GtkType.BUTTON_BOX_STYLE = g.GType(C.gtk_button_box_style_get_type())
	GtkType.DELETE_TYPE = g.GType(C.gtk_delete_type_get_type())
	GtkType.DIRECTION_TYPE = g.GType(C.gtk_direction_type_get_type())
	GtkType.EXPANDER_STYLE = g.GType(C.gtk_expander_style_get_type())
	GtkType.ICON_SIZE = g.GType(C.gtk_icon_size_get_type())
	GtkType.SENSITIVITY_TYPE = g.GType(C.gtk_sensitivity_type_get_type())
	GtkType.TEXT_DIRECTION = g.GType(C.gtk_text_direction_get_type())
	GtkType.JUSTIFICATION = g.GType(C.gtk_justification_get_type())
	GtkType.MENU_DIRECTION_TYPE = g.GType(C.gtk_menu_direction_type_get_type())
	GtkType.MESSAGE_TYPE = g.GType(C.gtk_message_type_get_type())
	GtkType.MOVEMENT_STEP = g.GType(C.gtk_movement_step_get_type())
	GtkType.SCROLL_STEP = g.GType(C.gtk_scroll_step_get_type())
	GtkType.ORIENTATION = g.GType(C.gtk_orientation_get_type())
	GtkType.CORNER_TYPE = g.GType(C.gtk_corner_type_get_type())
	GtkType.PACK_TYPE = g.GType(C.gtk_pack_type_get_type())
	GtkType.PATH_PRIORITY_TYPE = g.GType(C.gtk_path_priority_type_get_type())
	GtkType.PATH_TYPE = g.GType(C.gtk_path_type_get_type())
	GtkType.POLICY_TYPE = g.GType(C.gtk_policy_type_get_type())
	GtkType.POSITION_TYPE = g.GType(C.gtk_position_type_get_type())
	GtkType.RELIEF_STYLE = g.GType(C.gtk_relief_style_get_type())
	GtkType.RESIZE_MODE = g.GType(C.gtk_resize_mode_get_type())
	GtkType.SCROLL_TYPE = g.GType(C.gtk_scroll_type_get_type())
	GtkType.SELECTION_MODE = g.GType(C.gtk_selection_mode_get_type())
	GtkType.SHADOW_TYPE = g.GType(C.gtk_shadow_type_get_type())
	GtkType.STATE_TYPE = g.GType(C.gtk_state_type_get_type())
	GtkType.TOOLBAR_STYLE = g.GType(C.gtk_toolbar_style_get_type())
	GtkType.WINDOW_POSITION = g.GType(C.gtk_window_position_get_type())
	GtkType.WINDOW_TYPE = g.GType(C.gtk_window_type_get_type())
	GtkType.WRAP_MODE = g.GType(C.gtk_wrap_mode_get_type())
	GtkType.SORT_TYPE = g.GType(C.gtk_sort_type_get_type())
	GtkType.IM_PREEDIT_STYLE = g.GType(C.gtk_im_preedit_style_get_type())
	GtkType.IM_STATUS_STYLE = g.GType(C.gtk_im_status_style_get_type())
	GtkType.PACK_DIRECTION = g.GType(C.gtk_pack_direction_get_type())
	GtkType.PRINT_PAGES = g.GType(C.gtk_print_pages_get_type())
	GtkType.PAGE_SET = g.GType(C.gtk_page_set_get_type())
	GtkType.NUMBER_UP_LAYOUT = g.GType(C.gtk_number_up_layout_get_type())
	GtkType.PAGE_ORIENTATION = g.GType(C.gtk_page_orientation_get_type())
	GtkType.PRINT_QUALITY = g.GType(C.gtk_print_quality_get_type())
	GtkType.PRINT_DUPLEX = g.GType(C.gtk_print_duplex_get_type())
	GtkType.UNIT = g.GType(C.gtk_unit_get_type())
	GtkType.TREE_VIEW_GRID_LINES = g.GType(C.gtk_tree_view_grid_lines_get_type())
	GtkType.DRAG_RESULT = g.GType(C.gtk_drag_result_get_type())
	GtkType.SIZE_REQUEST_MODE = g.GType(C.gtk_size_request_mode_get_type())
	GtkType.SCROLLABLE_POLICY = g.GType(C.gtk_scrollable_policy_get_type())
	GtkType.STATE_FLAGS = g.GType(C.gtk_state_flags_get_type())
	GtkType.REGION_FLAGS = g.GType(C.gtk_region_flags_get_type())
	GtkType.JUNCTION_SIDES = g.GType(C.gtk_junction_sides_get_type())
	GtkType.BORDER_STYLE = g.GType(C.gtk_border_style_get_type())
	GtkType.FILE_CHOOSER_ACTION = g.GType(C.gtk_file_chooser_action_get_type())
	GtkType.FILE_CHOOSER_CONFIRMATION = g.GType(C.gtk_file_chooser_confirmation_get_type())
	GtkType.FILE_CHOOSER_ERROR = g.GType(C.gtk_file_chooser_error_get_type())
	GtkType.FILE_FILTER_FLAGS = g.GType(C.gtk_file_filter_flags_get_type())
	GtkType.ICON_LOOKUP_FLAGS = g.GType(C.gtk_icon_lookup_flags_get_type())
	GtkType.ICON_THEME_ERROR = g.GType(C.gtk_icon_theme_error_get_type())
	GtkType.ICON_VIEW_DROP_POSITION = g.GType(C.gtk_icon_view_drop_position_get_type())
	GtkType.IMAGE_TYPE = g.GType(C.gtk_image_type_get_type())
	GtkType.BUTTONS_TYPE = g.GType(C.gtk_buttons_type_get_type())
	GtkType.NOTEBOOK_TAB = g.GType(C.gtk_notebook_tab_get_type())
	GtkType.PRINT_STATUS = g.GType(C.gtk_print_status_get_type())
	GtkType.PRINT_OPERATION_RESULT = g.GType(C.gtk_print_operation_result_get_type())
	GtkType.PRINT_OPERATION_ACTION = g.GType(C.gtk_print_operation_action_get_type())
	GtkType.PRINT_ERROR = g.GType(C.gtk_print_error_get_type())
	GtkType.RC_FLAGS = g.GType(C.gtk_rc_flags_get_type())
	GtkType.RC_TOKEN_TYPE = g.GType(C.gtk_rc_token_type_get_type())
	GtkType.RECENT_SORT_TYPE = g.GType(C.gtk_recent_sort_type_get_type())
	GtkType.RECENT_CHOOSER_ERROR = g.GType(C.gtk_recent_chooser_error_get_type())
	GtkType.RECENT_FILTER_FLAGS = g.GType(C.gtk_recent_filter_flags_get_type())
	GtkType.RECENT_MANAGER_ERROR = g.GType(C.gtk_recent_manager_error_get_type())
	GtkType.SIZE_GROUP_MODE = g.GType(C.gtk_size_group_mode_get_type())
	GtkType.SPIN_BUTTON_UPDATE_POLICY = g.GType(C.gtk_spin_button_update_policy_get_type())
	GtkType.SPIN_TYPE = g.GType(C.gtk_spin_type_get_type())
	GtkType.TEXT_BUFFER_TARGET_INFO = g.GType(C.gtk_text_buffer_target_info_get_type())
	GtkType.TEXT_SEARCH_FLAGS = g.GType(C.gtk_text_search_flags_get_type())
	GtkType.TEXT_WINDOW_TYPE = g.GType(C.gtk_text_window_type_get_type())
	GtkType.TOOLBAR_SPACE_STYLE = g.GType(C.gtk_toolbar_space_style_get_type())
	GtkType.TOOL_PALETTE_DRAG_TARGETS = g.GType(C.gtk_tool_palette_drag_targets_get_type())
	GtkType.TREE_MODEL_FLAGS = g.GType(C.gtk_tree_model_flags_get_type())
	GtkType.TREE_VIEW_DROP_POSITION = g.GType(C.gtk_tree_view_drop_position_get_type())
	GtkType.TREE_VIEW_COLUMN_SIZING = g.GType(C.gtk_tree_view_column_sizing_get_type())
	GtkType.UI_MANAGER_ITEM_TYPE = g.GType(C.gtk_ui_manager_item_type_get_type())
	GtkType.WIDGET_HELP_TYPE = g.GType(C.gtk_widget_help_type_get_type())
    // End Initialize GtkType struct }}}

    // Other Gtk flags and constants {{{
	// Initialize GtkShadowType
	GtkShadow.NONE = 0
	GtkShadow.IN = 1
	GtkShadow.OUT = 2
	GtkShadow.ETCHED_IN = 3
	GtkShadow.ETCHED_OUT = 4

	// Initialize GtkStock
	GtkStock.ABOUT = "gtk-about"
	GtkStock.ADD = "gtk-add"
	GtkStock.APPLY = "gtk-apply"
	GtkStock.BOLD = "gtk-bold"
	GtkStock.CANCEL = "gtk-cancel"
	GtkStock.CAPS_LOCK_WARNING = "gtk-caps-lock-warning"
	GtkStock.CDROM = "gtk-cdrom"
	GtkStock.CLEAR = "gtk-clear"
	GtkStock.CLOSE = "gtk-close"
	GtkStock.COLOR_PICKER = "gtk-color-picker"
	GtkStock.CONNECT = "gtk-connect"
	GtkStock.CONVERT = "gtk-convert"
	GtkStock.COPY = "gtk-copy"
	GtkStock.CUT = "gtk-cut"
	GtkStock.DELETE = "gtk-delete"
	GtkStock.DIALOG_AUTHENTICATION = "gtk-dialog-authentication"
	GtkStock.DIALOG_INFO = "gtk-dialog-info"
	GtkStock.DIALOG_WARNING = "gtk-dialog-warning"
	GtkStock.DIALOG_ERROR = "gtk-dialog-error"
	GtkStock.DIALOG_QUESTION = "gtk-dialog-question"
	GtkStock.DIRECTORY = "gtk-directory"
	GtkStock.DISCARD = "gtk-discard"
	GtkStock.DISCONNECT = "gtk-disconnect"
	GtkStock.DND = "gtk-dnd"
	GtkStock.DND_MULTIPLE = "gtk-dnd-multiple"
	GtkStock.EDIT = "gtk-edit"
	GtkStock.EXECUTE = "gtk-execute"
	GtkStock.FILE = "gtk-file"
	GtkStock.FIND = "gtk-find"
	GtkStock.FIND_AND_REPLACE = "gtk-find-and-replace"
	GtkStock.FLOPPY = "gtk-floppy"
	GtkStock.FULLSCREEN = "gtk-fullscreen"
	GtkStock.GOTO_BOTTOM = "gtk-goto-bottom"
	GtkStock.GOTO_FIRST = "gtk-goto-first"
	GtkStock.GOTO_LAST = "gtk-goto-last"
	GtkStock.GOTO_TOP = "gtk-goto-top"
	GtkStock.GO_BACK = "gtk-go-back"
	GtkStock.GO_DOWN = "gtk-go-down"
	GtkStock.GO_FORWARD = "gtk-go-forward"
	GtkStock.GO_UP = "gtk-go-up"
	GtkStock.HARDDISK = "gtk-harddisk"
	GtkStock.HELP = "gtk-help"
	GtkStock.HOME = "gtk-home"
	GtkStock.INDEX = "gtk-index"
	GtkStock.INDENT = "gtk-indent"
	GtkStock.INFO = "gtk-info"
	GtkStock.ITALIC = "gtk-italic"
	GtkStock.JUMP_TO = "gtk-jump-to"
	GtkStock.JUSTIFY_CENTER = "gtk-justify-center"
	GtkStock.JUSTIFY_FILL = "gtk-justify-fill"
	GtkStock.JUSTIFY_LEFT = "gtk-justify-left"
	GtkStock.JUSTIFY_RIGHT = "gtk-justify-right"
	GtkStock.LEAVE_FULLSCREEN = "gtk-leave-fullscreen"
	GtkStock.MISSING_IMAGE = "gtk-missing-image"
	GtkStock.MEDIA_FORWARD = "gtk-media-forward"
	GtkStock.MEDIA_NEXT = "gtk-media-next"
	GtkStock.MEDIA_PAUSE = "gtk-media-pause"
	GtkStock.MEDIA_PLAY = "gtk-media-play"
	GtkStock.MEDIA_PREVIOUS = "gtk-media-previous"
	GtkStock.MEDIA_RECORD = "gtk-media-record"
	GtkStock.MEDIA_REWIND = "gtk-media-rewind"
	GtkStock.MEDIA_STOP = "gtk-media-stop"
	GtkStock.NETWORK = "gtk-network"
	GtkStock.NEW = "gtk-new"
	GtkStock.NO = "gtk-no"
	GtkStock.OK = "gtk-ok"
	GtkStock.OPEN = "gtk-open"
	GtkStock.ORIENTATION_PORTRAIT = "gtk-orientation-portrait"
	GtkStock.ORIENTATION_LANDSCAPE = "gtk-orientation-landscape"
	GtkStock.ORIENTATION_REVERSE_LANDSCAPE = "gtk-orientation-reverse-landscape"
	GtkStock.ORIENTATION_REVERSE_PORTRAIT = "gtk-orientation-reverse-portrait"
	GtkStock.PAGE_SETUP = "gtk-page-setup"
	GtkStock.PASTE = "gtk-paste"
	GtkStock.PREFERENCES = "gtk-preferences"
	GtkStock.PRINT = "gtk-print"
	GtkStock.PRINT_ERROR = "gtk-print-error"
	GtkStock.PRINT_PAUSED = "gtk-print-paused"
	GtkStock.PRINT_PREVIEW = "gtk-print-preview"
	GtkStock.PRINT_REPORT = "gtk-print-report"
	GtkStock.PRINT_WARNING = "gtk-print-warning"
	GtkStock.PROPERTIES = "gtk-properties"
	GtkStock.QUIT = "gtk-quit"
	GtkStock.REDO = "gtk-redo"
	GtkStock.REFRESH = "gtk-refresh"
	GtkStock.REMOVE = "gtk-remove"
	GtkStock.REVERT_TO_SAVED = "gtk-revert-to-saved"
	GtkStock.SAVE = "gtk-save"
	GtkStock.SAVE_AS = "gtk-save-as"
	GtkStock.SELECT_ALL = "gtk-select-all"
	GtkStock.SELECT_COLOR = "gtk-select-color"
	GtkStock.SELECT_FONT = "gtk-select-font"
	GtkStock.SORT_ASCENDING = "gtk-sort-ascending"
	GtkStock.SORT_DESCENDING = "gtk-sort-descending"
	GtkStock.SPELL_CHECK = "gtk-spell-check"
	GtkStock.STOP = "gtk-stop"
	GtkStock.STRIKETHROUGH = "gtk-strikethrough"
	GtkStock.UNDELETE = "gtk-undelete"
	GtkStock.UNDERLINE = "gtk-underline"
	GtkStock.UNDO = "gtk-undo"
	GtkStock.UNINDENT = "gtk-unindent"
	GtkStock.YES = "gtk-yes"
	GtkStock.ZOOM_100 = "gtk-zoom-100"
	GtkStock.ZOOM_FIT = "gtk-zoom-fit"
	GtkStock.ZOOM_IN = "gtk-zoom-in"
	GtkStock.ZOOM_OUT = "gtk-zoom-out"

	// Initialize GtkPosition
	GtkPosition.LEFT = 0
	GtkPosition.RIGHT = 1
	GtkPosition.TOP = 2
	GtkPosition.BOTTOM = 3

	// Initialize GtkDialogFlags
	GtkDialogFlags.MODAL = 1 << 0
	GtkDialogFlags.DESTROY_WITH_PARENT = 1 << 1

	// Initialize GtkResponse
	GtkResponse.NONE = -1
	GtkResponse.REJECT = -2
	GtkResponse.ACCEPT = -3
	GtkResponse.DELETE_EVENT = -4
	GtkResponse.OK = -5
	GtkResponse.CANCEL = -6
	GtkResponse.CLOSE = -7
	GtkResponse.YES = -8
	GtkResponse.NO = -9
	GtkResponse.APPLY = -10
	GtkResponse.HELP = -11

	// Initialize GtkMessage
	GtkMessage.INFO = 0
	GtkMessage.WARNING = 1
	GtkMessage.QUESTION = 2
	GtkMessage.ERROR = 3
	GtkMessage.OTHER = 4

	// Initialize GtkButtons
	GtkButtons.NONE = 0
	GtkButtons.OK = 1
	GtkButtons.CLOSE = 2
	GtkButtons.CANCEL = 3
	GtkButtons.YES_NO = 4
	GtkButtons.OK_CANCEL = 5

	// Initialize GtkOrientation
	GtkOrientation.HORIZONTAL = 0
	GtkOrientation.VERTICAL = 1

	// Initialize GtkPackType
	GtkPack.PACK_START = 0
	GtkPack.PACK_END = 1

	// Initialize GtkButtonBoxStyle
	GtkButtonBoxStyle.SPREAD = 1
	GtkButtonBoxStyle.EDGE = 2
	GtkButtonBoxStyle.START = 3
	GtkButtonBoxStyle.END = 4
	GtkButtonBoxStyle.CENTER = 5

	// Initialize GtkIconSize
	GtkIconSize.INVALID = 0
	GtkIconSize.MENU = 1
	GtkIconSize.SMALL_TOOLBAR = 2
	GtkIconSize.LARGE_TOOLBAR = 3
	GtkIconSize.BUTTON = 4
	GtkIconSize.DND = 5
	GtkIconSize.DIALOG = 6

	// Initialize GApplicationFlags
	GApplicationFlags.FLAGS_NONE = 0
	GApplicationFlags.IS_SERVICE = 1 << 0
	GApplicationFlags.IS_LAUNCHER = 1 << 1
	GApplicationFlags.HANDLES_OPEN = 1 << 2
	GApplicationFlags.HANDLES_COMMAND_LINE = 1 << 3
	GApplicationFlags.SEND_ENVIRONMENT = 1 << 4
	GApplicationFlags.NON_UNIQUE = 1 << 5

	// Initialize GtkJustification
	GtkJustification.LEFT = 0
	GtkJustification.RIGHT = 1
	GtkJustification.CENTER = 2
	GtkJustification.FILL = 3

	// Initialize GtkWindowType
	GtkWindowType.TOPLEVEL = 0
	GtkWindowType.POPUP = 1

	// Initialize GtkSensitivity
	GtkSensitivity.AUTO = 0
	GtkSensitivity.ON = 1
	GtkSensitivity.OFF = 2

	// Initialize GtkPolicy
	GtkPolicy.ALWAYS = 0
	GtkPolicy.AUTOMATIC = 1
	GtkPolicy.NEVER = 2

	// Initialize GtkCorner
	GtkCorner.TOP_LEFT = 0
	GtkCorner.BOTTOM_LEFT = 1
	GtkCorner.TOP_RIGHT = 2
	GtkCorner.BOTTOM_RIGHT = 3

	// Initialize GtkTreeModelFlags
	GtkTreeModelFlags.ITERS_PERSIST = 1 << 0
	GtkTreeModelFlags.LIST_ONLY = 1 << 1

	// Initialize GtkCellRendererState
	GtkCellRendererState.SELECTED = 1 << 0
	GtkCellRendererState.PRELIT = 1 << 1
	GtkCellRendererState.INSENSITIVE = 1 << 2
	GtkCellRendererState.SORTED = 1 << 3
	GtkCellRendererState.FOCUSED = 1 << 4

	// Initialize GtkCellRendererMode
	GtkCellRendererMode.INERT = 0
	GtkCellRendererMode.ACTIVATABLE = 1
	GtkCellRendererMode.EDITABLE = 2

	// Initialize GtkSizeRequestMode
	GtkSizeRequestMode.HEIGHT_FOR_WIDTH = 0
	GtkSizeRequestMode.WIDTH_FOR_HEIGHT = 1
	GtkSizeRequestMode.CONSTANT_SIZE = 2

	// Initialize GtkCellRendererAccelMode
	GtkCellRendererAccelMode.GTK = 0
	GtkCellRendererAccelMode.OTHER = 1

	// Initialize GtkTreeViewColumnSizing
	GtkTreeViewColumnSizing.GROW_ONLY = 0
	GtkTreeViewColumnSizing.AUTOSIZE = 1
	GtkTreeViewColumnSizing.FIXED = 2

	// Initialize GtkSortType
	GtkSortType.ASCENDING = 0
	GtkSortType.DESCENDING = 1

	// Initialize GtkSelectionMode
	GtkSelectionMode.NONE = 0
	GtkSelectionMode.SINGLE = 1
	GtkSelectionMode.BROWSE = 2
	GtkSelectionMode.MULTIPLE = 3

	// Initialize GtkResizeMode
	GtkResizeMode.PARENT = 0
	GtkResizeMode.QUEUE = 1

	// Initialize GtkAssistantPage
	GtkAssistantPage.CONTENT = 0
	GtkAssistantPage.INTRO = 1
	GtkAssistantPage.CONFIRM = 2
	GtkAssistantPage.SUMMARY = 3
	GtkAssistantPage.PROGRESS = 4
	GtkAssistantPage.CUSTOM = 5

	// Initialize GtkAlign
	GtkAlign.FILL = 0
	GtkAlign.START = 1
	GtkAlign.END = 2
	GtkAlign.CENTER = 3

	// Initialize GtkTextSearchFlags
	GtkTextSearchFlags.VISIBLE_ONLY = 1 << 0
	GtkTextSearchFlags.TEXT_ONLY = 1 << 1
	GtkTextSearchFlags.CASE_INSENSITIVE = 1 << 2

	// Initialize GtkTextWindowType
	GtkTextWindowType.PRIVATE = 0
	GtkTextWindowType.WIDGET = 1
	GtkTextWindowType.TEXT = 2
	GtkTextWindowType.LEFT = 3
	GtkTextWindowType.RIGHT = 4
	GtkTextWindowType.TOP = 5
	GtkTextWindowType.BOTTOM = 6

	// Initialize GtkWrapMode
	GtkWrapMode.NONE = 0
	GtkWrapMode.CHAR = 1
	GtkWrapMode.WORD = 2
	GtkWrapMode.WORD_CHAR = 3

	// Initialize GtkTextDirection
	GtkTextDirection.NONE = 0
	GtkTextDirection.LTR = 1
	GtkTextDirection.RTL = 2

	// Initialize GtkAccelFlags
	GtkAccelFlags.VISIBLE = 1 << 0
	GtkAccelFlags.LOCKED = 1 << 1
	GtkAccelFlags.MASK = 0x07

	// Initialize GtkMenuDirection
	GtkMenuDirection.PARENT = 0
	GtkMenuDirection.CHILD = 1
	GtkMenuDirection.NEXT = 2
	GtkMenuDirection.PREV = 3

	// Initialize GtkPackDirection
	GtkPackDirection.LTR = 0
	GtkPackDirection.RTL = 1
	GtkPackDirection.TTB = 2
	GtkPackDirection.BTT = 3

	// Initialize GtkSpinButtonUpdatePolicy
	GtkSpinButtonUpdatePolicy.ALWAYS = 0
	GtkSpinButtonUpdatePolicy.IF_VALID = 1

	// Initialize GtkSpinType
	GtkSpinType.STEP_FORWARD = 0
	GtkSpinType.STEP_BACKWARD = 1
	GtkSpinType.PAGE_FORWARD = 2
	GtkSpinType.PAGE_BACKWARD = 3
	GtkSpinType.HOME = 4
	GtkSpinType.END = 5
	GtkSpinType.USER_DEFINED = 6

	// Initialize GtkTreeViewGridLines
	GtkTreeViewGridLines.NONE = 0
	GtkTreeViewGridLines.HORIZONTAL = 1
	GtkTreeViewGridLines.VERTICAL = 2
	GtkTreeViewGridLines.BOTH = 3

	// Initialize GtkEntryIconPosition
	GtkEntryIconPosition.PRIMARY = 0
	GtkEntryIconPosition.SECONDARY = 1

	// Initialize GtkImageType
	GtkImageType.EMPTY = 0
	GtkImageType.PIXBUF = 1
	GtkImageType.STOCK = 2
	GtkImageType.ICON_SET = 3
	GtkImageType.ANIMATION = 4
	GtkImageType.ICON_NAME = 5
	GtkImageType.GICON = 6
    // End Other Gtk flags and constants }}}

	// Register GtkAccelGroup
	gobject.RegisterCType(GtkType.ACCEL_GROUP, newAccelGroupFromNative)
	gobject.RegisterGoType(GtkType.ACCEL_GROUP, nativeFromAccelGroup)

	// Register GtkClipboard type
	gobject.RegisterCType(GtkType.CLIPBOARD, newClipboardFromNative)

	// Register GtkWidget type
	gobject.RegisterCType(GtkType.WIDGET, newWidgetFromNative)

	// Register GtkBin type
	gobject.RegisterCType(GtkType.BIN, newBinFromNative)
	gobject.RegisterGoType(GtkType.BIN, nativeFromBin)

	// Register GtkApplicaton type
	gobject.RegisterCType(GtkType.APPLICATION, appFromNative)
	gobject.RegisterGoType(GtkType.APPLICATION, nativeFromApp)

	// Register GtkWindow type
	gobject.RegisterCType(GtkType.WINDOW, newWindowFromNative)
	gobject.RegisterGoType(GtkType.WINDOW, nativeFromWindow)

	// Register GtkInvisible type
	gobject.RegisterCType(GtkType.INVISIBLE, newInvisibleFromNative)
	gobject.RegisterGoType(GtkType.INVISIBLE, nativeFromInvisible)

	// Register GtkAssistant type
	gobject.RegisterCType(GtkType.ASSISTANT, newAssistantFromNative)
	gobject.RegisterGoType(GtkType.ASSISTANT, nativeFromAssistant)

	// Register GtkBox type
	gobject.RegisterCType(GtkType.BOX, newBoxFromNative)
	gobject.RegisterGoType(GtkType.BOX, nativeFromBox)

	// Regiseter GtkPaned type
	gobject.RegisterCType(GtkType.PANED, newPanedFromNative)
	gobject.RegisterGoType(GtkType.PANED, nativeFromPaned)

	// Register GtkButtonBox type
	gobject.RegisterCType(GtkType.BUTTON_BOX, newButtonBoxFromNative)
	gobject.RegisterGoType(GtkType.BUTTON_BOX, nativeFromButtonBox)

	// Register GtkFrame type
	gobject.RegisterCType(GtkType.FRAME, newFrameFromNative)
	gobject.RegisterGoType(GtkType.FRAME, nativeFromFrame)

	// Register GtkGrid type
	gobject.RegisterCType(GtkType.GRID, newGridFromNative)
	gobject.RegisterGoType(GtkType.GRID, nativeFromGrid)

	// Register GtkLabel type
	gobject.RegisterCType(GtkType.LABEL, newLabelFromNative)
	gobject.RegisterGoType(GtkType.LABEL, nativeFromLabel)

	// Register GtkProgressBar type
	gobject.RegisterCType(GtkType.PROGRESS_BAR, newProgressBarFromNative)
	gobject.RegisterGoType(GtkType.PROGRESS_BAR, nativeFromProgressBar)

	// Register GtkStatusbar type
	gobject.RegisterCType(GtkType.STATUSBAR, newStatusbarFromNative)
	gobject.RegisterGoType(GtkType.STATUSBAR, nativeFromStatusbar)

    // Register GtkSpinner
    gobject.RegisterCType(GtkType.SPINNER, newSpinnerFromNative)
    gobject.RegisterGoType(GtkType.SPINNER, nativeFromSpinner)

    // Register GtkInfoBar
    gobject.RegisterCType(GtkType.INFO_BAR, newInfoBarFromNative)
    gobject.RegisterGoType(GtkType.INFO_BAR, nativeFromInfoBar)

	// Register GtkImage type
	gobject.RegisterCType(GtkType.IMAGE, newImageFromNative)
	gobject.RegisterGoType(GtkType.IMAGE, nativeFromImage)

	// Register GtkButton type
	gobject.RegisterCType(GtkType.BUTTON, newButtonFromNative)
	gobject.RegisterGoType(GtkType.BUTTON, nativeFromButton)

	//Register GtkToggleButton type
	gobject.RegisterCType(GtkType.TOGGLE_BUTTON, newToggleButtonFromNative)
	gobject.RegisterGoType(GtkType.TOGGLE_BUTTON, nativeFromToggleButton)

	// Register GtkCheckButton type
	gobject.RegisterCType(GtkType.CHECK_BUTTON, newCheckButtonFromNative)
	gobject.RegisterGoType(GtkType.CHECK_BUTTON, nativeFromCheckButton)

	// Register GtkRadioButton type
	gobject.RegisterCType(GtkType.RADIO_BUTTON, newRadioButtonFromNative)
	gobject.RegisterGoType(GtkType.RADIO_BUTTON, nativeFromRadioButton)

    gobject.RegisterCType(GtkType.SWITCH, newSwitchFromNative)
    gobject.RegisterGoType(GtkType.SWITCH, nativeFromSwitch)

	// Register GtkEntryBuffer type
	gobject.RegisterCType(GtkType.ENTRY_BUFFER, newEntryBufferFromNative)
	gobject.RegisterGoType(GtkType.ENTRY_BUFFER, nativeFromEntryBuffer)

	// Register GtkEntry type
	gobject.RegisterCType(GtkType.ENTRY, newEntryFromNative)
	gobject.RegisterGoType(GtkType.ENTRY, nativeFromEntry)

	// Register GtkEntryCompletion type
	gobject.RegisterCType(GtkType.ENTRY_COMPLETION, newEntryCompletionFromNative)
	gobject.RegisterGoType(GtkType.ENTRY_COMPLETION, nativeFromEntryCompletion)

	// Register GtkScale type
	gobject.RegisterCType(GtkType.SCALE, newScaleFromNative)
	gobject.RegisterGoType(GtkType.SCALE, nativeFromScale)

	// Register GtkSpinButton type
	gobject.RegisterCType(GtkType.SPIN_BUTTON, newSpinButtonFromNative)
	gobject.RegisterGoType(GtkType.SPIN_BUTTON, nativeFromSpinButton)

	// Register GtkDialog type
	gobject.RegisterCType(GtkType.DIALOG, newDialogFromNative)
	gobject.RegisterGoType(GtkType.DIALOG, nativeFromDialog)

	// Register GtkMessageDialog type
	gobject.RegisterCType(GtkType.MESSAGE_DIALOG, newMessageDialogFromNative)
	gobject.RegisterGoType(GtkType.MESSAGE_DIALOG, nativeFromMessageDialog)

	// Register GtkSeparator type
	gobject.RegisterCType(GtkType.SEPARATOR, newSeparatorFromNative)
	gobject.RegisterGoType(GtkType.SEPARATOR, nativeFromSeparator)

	// Register GtkAdjustment type
	gobject.RegisterCType(GtkType.ADJUSTMENT, newAdjustmentFromNative)
	gobject.RegisterGoType(GtkType.ADJUSTMENT, nativeFromAdjustment)

	// Register GtkRange type
	gobject.RegisterCType(GtkType.RANGE, newRangeFromNative)
	gobject.RegisterGoType(GtkType.RANGE, nativeFromRange)

	// Register GtkScrollbar type
	gobject.RegisterCType(GtkType.SCROLLBAR, newScrollbarFromNative)
	gobject.RegisterGoType(GtkType.SCROLLBAR, nativeFromScrollbar)

	// Register GtkScrolledWindow type
	gobject.RegisterCType(GtkType.SCROLLED_WINDOW, newScrolledWindowFromNative)
	gobject.RegisterGoType(GtkType.SCROLLED_WINDOW, nativeFromScrolledWindow)

	// Register GtkTextTag type
	gobject.RegisterCType(GtkType.TEXT_TAG, newTextTagFromNative)
	gobject.RegisterGoType(GtkType.TEXT_TAG, nativeFromTextTag)

	// Register GtkTextTagTable type
	gobject.RegisterCType(GtkType.TEXT_TAG_TABLE, newTextTagTableFromNative)
	gobject.RegisterGoType(GtkType.TEXT_TAG_TABLE, nativeFromTextTagTable)

	// Register GtkTextIter type
	gobject.RegisterCType(GtkType.TEXT_ITER, newTextIterFromNative)
	gobject.RegisterGoType(GtkType.TEXT_ITER, nativeFromTextIter)

	// Register GtkTextBuffer type
	gobject.RegisterCType(GtkType.TEXT_BUFFER, newTextBufferFromNative)
	gobject.RegisterGoType(GtkType.TEXT_BUFFER, nativeFromTextBuffer)

	// Register GtkTextChildAnchor
	gobject.RegisterCType(GtkType.TEXT_CHILD_ANCHOR, newTextChildAnchorFromNative)
	gobject.RegisterGoType(GtkType.TEXT_CHILD_ANCHOR, nativeFromTextChildAnchor)

	// Register GtkTextView type
	gobject.RegisterCType(GtkType.TEXT_VIEW, newTextViewFromNative)
	gobject.RegisterGoType(GtkType.TEXT_VIEW, nativeFromTextView)

	// Register TreePath type
	gobject.RegisterCType(GtkType.TREE_PATH, newTreePathFromNative)
	// Register GtkTreeModel type
	gobject.RegisterCType(GtkType.TREE_MODEL, newTreeModelFromNative)
	gobject.RegisterGoType(GtkType.TREE_MODEL, nativeFromTreeModel)

	// Register GtkListStore type
	gobject.RegisterCType(GtkType.LIST_STORE, newListStoreFromNative)
	gobject.RegisterGoType(GtkType.LIST_STORE, nativeFromListStore)

	// Register GtkTreeStore type
	gobject.RegisterCType(GtkType.TREE_STORE, newTreeStoreFromNative)
	gobject.RegisterGoType(GtkType.TREE_STORE, nativeFromTreeStore)

	// Register GtkCellArea type
	gobject.RegisterCType(GtkType.CELL_AREA, newCellAreaFromNative)
	gobject.RegisterGoType(GtkType.CELL_AREA, nativeFromCellArea)

	// Register GtkCellAreaContext type
	gobject.RegisterCType(GtkType.CELL_AREA_CONTEXT, newCellAreaContextFromNative)
	gobject.RegisterGoType(GtkType.CELL_AREA_CONTEXT, nativeFromCellAreaContext)

	// Register GtkCellLayout interface type
	gobject.RegisterCType(GtkType.CELL_LAYOUT, newCellLayoutFromNative)
	gobject.RegisterGoType(GtkType.CELL_LAYOUT, nativeFromCellLayout)

	// Register GtkCellView type
	gobject.RegisterCType(GtkType.CELL_VIEW, newCellViewFromNative)
	gobject.RegisterGoType(GtkType.CELL_VIEW, nativeFromCellView)

	// Register GtkCellRenderer type
	gobject.RegisterCType(GtkType.CELL_RENDERER, newCellRendererFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER, nativeFromCellRenderer)

	// Register GtkCellRendererText type
	gobject.RegisterCType(GtkType.CELL_RENDERER_TEXT, newCellRendererTextFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_TEXT, nativeFromCellRendererText)

	// Register GtkCellRendererProgress type
	gobject.RegisterCType(GtkType.CELL_RENDERER_PROGRESS, newCellRendererProgressFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_PROGRESS, nativeFromCellRendererProgress)

	// Register GtkCellRendererSpinner type
	gobject.RegisterCType(GtkType.CELL_RENDERER_SPINNER, newCellRendererSpinnerFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_SPINNER, nativeFromCellRendererSpinner)

	// Register GtkCellRendererToggle type
	gobject.RegisterCType(GtkType.CELL_RENDERER_TOGGLE, newCellRendererToggleFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_TOGGLE, nativeFromCellRendererToggle)

	// Register GtkCellRendererPixbuf type
	gobject.RegisterCType(GtkType.CELL_RENDERER_PIXBUF, newCellRendererPixbufFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_PIXBUF, nativeFromCellRendererPixbuf)

	// Register GtkCellRendererAccel type
	gobject.RegisterCType(GtkType.CELL_RENDERER_ACCEL, newCellRendererAccelFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_ACCEL, nativeFromCellRendererAccel)

	// Register GtkCellRendererCombo type
	gobject.RegisterCType(GtkType.CELL_RENDERER_COMBO, newCellRendererComboFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_COMBO, nativeFromCellRendererCombo)

	// Register GtkCellRendererSpin type
	gobject.RegisterCType(GtkType.CELL_RENDERER_SPIN, newCellRendererSpinFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_SPIN, nativeFromCellRendererSpin)

	// Register GtkTreeViewColumn type
	gobject.RegisterCType(GtkType.TREE_VIEW_COLUMN, newTreeViewColumnFromNative)
	gobject.RegisterGoType(GtkType.TREE_VIEW_COLUMN, nativeFromTreeViewColumn)

	// Register GtkTreeSelection type
	gobject.RegisterCType(GtkType.TREE_SELECTION, newTreeSelectionFromNative)
	gobject.RegisterGoType(GtkType.TREE_SELECTION, nativeFromTreeSelection)

	// Register GtkTreeView type
	gobject.RegisterCType(GtkType.TREE_VIEW, newTreeViewFromNative)
	gobject.RegisterGoType(GtkType.TREE_VIEW, nativeFromTreeView)

	// Register GtkComboBox type
	gobject.RegisterCType(GtkType.COMBO_BOX, newComboBoxFromNative)
	gobject.RegisterGoType(GtkType.COMBO_BOX, nativeFromComboBox)

	// Register GtkComboBoxText type
	gobject.RegisterCType(GtkType.COMBO_BOX_TEXT, newComboBoxTextFromNative)
	gobject.RegisterGoType(GtkType.COMBO_BOX_TEXT, nativeFromComboBoxText)

	// Register GtkMenuShell type
	gobject.RegisterCType(GtkType.MENU_SHELL, newMenuShellFromNative)
	gobject.RegisterGoType(GtkType.MENU_SHELL, nativeFromMenuShell)

	// Register GtkMenuDirection type
	gobject.RegisterCType(GtkType.MENU_DIRECTION_TYPE, newMenuShellFromNative)

	// Register GtkMenu type
	gobject.RegisterCType(GtkType.MENU, newMenuFromNative)
	gobject.RegisterGoType(GtkType.MENU, nativeFromMenu)

	// Register GtkMenuBar type
	gobject.RegisterCType(GtkType.MENU_BAR, newMenuBarFromNative)
	gobject.RegisterGoType(GtkType.MENU_BAR, nativeFromMenuBar)

	gobject.RegisterCType(GtkType.MENU_ITEM, newMenuItemFromNative)
	gobject.RegisterGoType(GtkType.MENU_ITEM, nativeFromMenuItem)

	// Register GtkCheckMenuItem type
	gobject.RegisterCType(GtkType.CHECK_MENU_ITEM, newCheckMenuItemFromNative)
	gobject.RegisterGoType(GtkType.CHECK_MENU_ITEM, nativeFromCheckMenuItem)

	// Register GtkRadioMenuItem type
	gobject.RegisterCType(GtkType.RADIO_MENU_ITEM, newRadioMenuItemFromNative)
	gobject.RegisterGoType(GtkType.RADIO_MENU_ITEM, nativeFromRadioMenuItem)

	// Register GtkImageMenuItem type
	gobject.RegisterCType(GtkType.IMAGE_MENU_ITEM, newImageMenuItemFromNative)
	gobject.RegisterGoType(GtkType.IMAGE_MENU_ITEM, nativeFromImageMenuItem)

	// Register GtkSeparatorMenuItem type
	gobject.RegisterCType(GtkType.SEPARATOR_MENU_ITEM, newSeparatorMenuItemFromNative)
	gobject.RegisterGoType(GtkType.SEPARATOR_MENU_ITEM, nativeFromSeparatorMenuItem)

	// Register GtkTearoffMenuitem type
	gobject.RegisterCType(GtkType.TEAROFF_MENU_ITEM, newTearoffMenuItemFromNative)
	gobject.RegisterGoType(GtkType.TEAROFF_MENU_ITEM, nativeFromTearoffMenuItem)

    // Register GtkColorbutton type
    gobject.RegisterCType(GtkType.COLOR_BUTTON, newColorButtonFromNative)
    gobject.RegisterGoType(GtkType.COLOR_BUTTON, nativeFromColorButton)

	// Register GtkNotebook type
	gobject.RegisterCType(GtkType.NOTEBOOK, newNotebookFromNative)
	gobject.RegisterGoType(GtkType.NOTEBOOK, nativeFromNotebook)

	// Register GtkOrientable interface type
	gobject.RegisterCType(GtkType.ORIENTABLE, newOrientableFromNative)
	gobject.RegisterGoType(GtkType.ORIENTABLE, nativeFromOrientable)

	// Register GtkAction type
	gobject.RegisterCType(GtkType.ACTION, newActionFromNative)
	gobject.RegisterGoType(GtkType.ACTION, nativeFromAction)
}

// End init function }}}
