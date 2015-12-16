package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRasterWindow struct {
	QPaintDeviceWindow
}

type QRasterWindow_ITF interface {
	QPaintDeviceWindow_ITF
	QRasterWindow_PTR() *QRasterWindow
}

func PointerFromQRasterWindow(ptr QRasterWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRasterWindow_PTR().Pointer()
	}
	return nil
}

func NewQRasterWindowFromPointer(ptr unsafe.Pointer) *QRasterWindow {
	var n = new(QRasterWindow)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRasterWindow_") {
		n.SetObjectName("QRasterWindow_" + qt.Identifier())
	}
	return n
}

func (ptr *QRasterWindow) QRasterWindow_PTR() *QRasterWindow {
	return ptr
}

func NewQRasterWindow(parent QWindow_ITF) *QRasterWindow {
	defer qt.Recovering("QRasterWindow::QRasterWindow")

	return NewQRasterWindowFromPointer(C.QRasterWindow_NewQRasterWindow(PointerFromQWindow(parent)))
}