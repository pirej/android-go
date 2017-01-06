// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Fri, 06 Jan 2017 23:23:59 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package egl

/*
#cgo LDFLAGS: -lEGL
#include <EGL/egl.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// GetError function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglGetError.xhtml
func GetError() Int {
	__ret := C.eglGetError()
	__v := (Int)(__ret)
	return __v
}

// GetDisplay function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglGetDisplay.xhtml
func GetDisplay(displayId NativeDisplayType) Display {
	cdisplayId, _ := (unsafe.Pointer)(displayId), cgoAllocsUnknown
	__ret := C.eglGetDisplay(cdisplayId)
	__v := (Display)(__ret)
	return __v
}

// Initialize function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglInitialize.xhtml
func Initialize(dpy Display, major *Int, minor *Int) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cmajor, _ := (*C.EGLint)(unsafe.Pointer(major)), cgoAllocsUnknown
	cminor, _ := (*C.EGLint)(unsafe.Pointer(minor)), cgoAllocsUnknown
	__ret := C.eglInitialize(cdpy, cmajor, cminor)
	__v := (Boolean)(__ret)
	return __v
}

// Terminate function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglTerminate.xhtml
func Terminate(dpy Display) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	__ret := C.eglTerminate(cdpy)
	__v := (Boolean)(__ret)
	return __v
}

// QueryString function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglQueryString.xhtml
func QueryString(dpy Display, name Int) string {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cname, _ := (C.EGLint)(name), cgoAllocsUnknown
	__ret := C.eglQueryString(cdpy, cname)
	__v := packPCharString(__ret)
	return __v
}

// GetConfigs function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglGetConfigs.xhtml
func GetConfigs(dpy Display, configs []Config, configSize Int, numConfig *Int) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cconfigs, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&configs)).Data)), cgoAllocsUnknown
	cconfigSize, _ := (C.EGLint)(configSize), cgoAllocsUnknown
	cnumConfig, _ := (*C.EGLint)(unsafe.Pointer(numConfig)), cgoAllocsUnknown
	__ret := C.eglGetConfigs(cdpy, cconfigs, cconfigSize, cnumConfig)
	__v := (Boolean)(__ret)
	return __v
}

// ChooseConfig function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglChooseConfig.xhtml
func ChooseConfig(dpy Display, attribList []Int, configs []Config, configSize Int, numConfig *Int) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cattribList, _ := (*C.EGLint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&attribList)).Data)), cgoAllocsUnknown
	cconfigs, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&configs)).Data)), cgoAllocsUnknown
	cconfigSize, _ := (C.EGLint)(configSize), cgoAllocsUnknown
	cnumConfig, _ := (*C.EGLint)(unsafe.Pointer(numConfig)), cgoAllocsUnknown
	__ret := C.eglChooseConfig(cdpy, cattribList, cconfigs, cconfigSize, cnumConfig)
	__v := (Boolean)(__ret)
	return __v
}

// GetConfigAttrib function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglGetConfigAttrib.xhtml
func GetConfigAttrib(dpy Display, config Config, attribute Int, value *Int) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cconfig, _ := (unsafe.Pointer)(config), cgoAllocsUnknown
	cattribute, _ := (C.EGLint)(attribute), cgoAllocsUnknown
	cvalue, _ := (*C.EGLint)(unsafe.Pointer(value)), cgoAllocsUnknown
	__ret := C.eglGetConfigAttrib(cdpy, cconfig, cattribute, cvalue)
	__v := (Boolean)(__ret)
	return __v
}

// CreateWindowSurface function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreateWindowSurface.xhtml
func CreateWindowSurface(dpy Display, config Config, win NativeWindowType, attribList []Int) Surface {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cconfig, _ := (unsafe.Pointer)(config), cgoAllocsUnknown
	cwin, _ := *(*C.EGLNativeWindowType)(unsafe.Pointer(&win)), cgoAllocsUnknown
	cattribList, _ := (*C.EGLint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&attribList)).Data)), cgoAllocsUnknown
	__ret := C.eglCreateWindowSurface(cdpy, cconfig, cwin, cattribList)
	__v := (Surface)(__ret)
	return __v
}

// CreatePbufferSurface function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreatePbufferSurface.xhtml
func CreatePbufferSurface(dpy Display, config Config, attribList []Int) Surface {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cconfig, _ := (unsafe.Pointer)(config), cgoAllocsUnknown
	cattribList, _ := (*C.EGLint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&attribList)).Data)), cgoAllocsUnknown
	__ret := C.eglCreatePbufferSurface(cdpy, cconfig, cattribList)
	__v := (Surface)(__ret)
	return __v
}

// CreatePixmapSurface function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreatePixmapSurface.xhtml
func CreatePixmapSurface(dpy Display, config Config, pixmap NativePixmapType, attribList []Int) Surface {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cconfig, _ := (unsafe.Pointer)(config), cgoAllocsUnknown
	cpixmap, _ := *(*C.EGLNativePixmapType)(unsafe.Pointer(&pixmap)), cgoAllocsUnknown
	cattribList, _ := (*C.EGLint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&attribList)).Data)), cgoAllocsUnknown
	__ret := C.eglCreatePixmapSurface(cdpy, cconfig, cpixmap, cattribList)
	__v := (Surface)(__ret)
	return __v
}

// DestroySurface function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglDestroySurface.xhtml
func DestroySurface(dpy Display, surface Surface) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	csurface, _ := (unsafe.Pointer)(surface), cgoAllocsUnknown
	__ret := C.eglDestroySurface(cdpy, csurface)
	__v := (Boolean)(__ret)
	return __v
}

// QuerySurface function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglQuerySurface.xhtml
func QuerySurface(dpy Display, surface Surface, attribute Int, value *Int) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	csurface, _ := (unsafe.Pointer)(surface), cgoAllocsUnknown
	cattribute, _ := (C.EGLint)(attribute), cgoAllocsUnknown
	cvalue, _ := (*C.EGLint)(unsafe.Pointer(value)), cgoAllocsUnknown
	__ret := C.eglQuerySurface(cdpy, csurface, cattribute, cvalue)
	__v := (Boolean)(__ret)
	return __v
}

// BindAPI function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglBindAPI.xhtml
func BindAPI(api Enum) Boolean {
	capi, _ := (C.EGLenum)(api), cgoAllocsUnknown
	__ret := C.eglBindAPI(capi)
	__v := (Boolean)(__ret)
	return __v
}

// QueryAPI function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglQueryAPI.xhtml
func QueryAPI() Enum {
	__ret := C.eglQueryAPI()
	__v := (Enum)(__ret)
	return __v
}

// WaitClient function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglWaitClient.xhtml
func WaitClient() Boolean {
	__ret := C.eglWaitClient()
	__v := (Boolean)(__ret)
	return __v
}

// ReleaseThread function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglReleaseThread.xhtml
func ReleaseThread() Boolean {
	__ret := C.eglReleaseThread()
	__v := (Boolean)(__ret)
	return __v
}

// CreatePbufferFromClientBuffer function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreatePbufferFromClientBuffer.xhtml
func CreatePbufferFromClientBuffer(dpy Display, buftype Enum, buffer ClientBuffer, config Config, attribList []Int) Surface {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cbuftype, _ := (C.EGLenum)(buftype), cgoAllocsUnknown
	cbuffer, _ := (unsafe.Pointer)(buffer), cgoAllocsUnknown
	cconfig, _ := (unsafe.Pointer)(config), cgoAllocsUnknown
	cattribList, _ := (*C.EGLint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&attribList)).Data)), cgoAllocsUnknown
	__ret := C.eglCreatePbufferFromClientBuffer(cdpy, cbuftype, cbuffer, cconfig, cattribList)
	__v := (Surface)(__ret)
	return __v
}

// SurfaceAttrib function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglSurfaceAttrib.xhtml
func SurfaceAttrib(dpy Display, surface Surface, attribute Int, value Int) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	csurface, _ := (unsafe.Pointer)(surface), cgoAllocsUnknown
	cattribute, _ := (C.EGLint)(attribute), cgoAllocsUnknown
	cvalue, _ := (C.EGLint)(value), cgoAllocsUnknown
	__ret := C.eglSurfaceAttrib(cdpy, csurface, cattribute, cvalue)
	__v := (Boolean)(__ret)
	return __v
}

// BindTexImage function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglBindTexImage.xhtml
func BindTexImage(dpy Display, surface Surface, buffer Int) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	csurface, _ := (unsafe.Pointer)(surface), cgoAllocsUnknown
	cbuffer, _ := (C.EGLint)(buffer), cgoAllocsUnknown
	__ret := C.eglBindTexImage(cdpy, csurface, cbuffer)
	__v := (Boolean)(__ret)
	return __v
}

// ReleaseTexImage function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglReleaseTexImage.xhtml
func ReleaseTexImage(dpy Display, surface Surface, buffer Int) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	csurface, _ := (unsafe.Pointer)(surface), cgoAllocsUnknown
	cbuffer, _ := (C.EGLint)(buffer), cgoAllocsUnknown
	__ret := C.eglReleaseTexImage(cdpy, csurface, cbuffer)
	__v := (Boolean)(__ret)
	return __v
}

// SwapInterval function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglSwapInterval.xhtml
func SwapInterval(dpy Display, interval Int) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cinterval, _ := (C.EGLint)(interval), cgoAllocsUnknown
	__ret := C.eglSwapInterval(cdpy, cinterval)
	__v := (Boolean)(__ret)
	return __v
}

// CreateContext function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreateContext.xhtml
func CreateContext(dpy Display, config Config, shareContext Context, attribList *Int) Context {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cconfig, _ := (unsafe.Pointer)(config), cgoAllocsUnknown
	cshareContext, _ := (unsafe.Pointer)(shareContext), cgoAllocsUnknown
	cattribList, _ := (*C.EGLint)(unsafe.Pointer(attribList)), cgoAllocsUnknown
	__ret := C.eglCreateContext(cdpy, cconfig, cshareContext, cattribList)
	__v := (Context)(__ret)
	return __v
}

// DestroyContext function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglDestroyContext.xhtml
func DestroyContext(dpy Display, ctx Context) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cctx, _ := (unsafe.Pointer)(ctx), cgoAllocsUnknown
	__ret := C.eglDestroyContext(cdpy, cctx)
	__v := (Boolean)(__ret)
	return __v
}

// MakeCurrent function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglMakeCurrent.xhtml
func MakeCurrent(dpy Display, draw Surface, read Surface, ctx Context) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cdraw, _ := (unsafe.Pointer)(draw), cgoAllocsUnknown
	cread, _ := (unsafe.Pointer)(read), cgoAllocsUnknown
	cctx, _ := (unsafe.Pointer)(ctx), cgoAllocsUnknown
	__ret := C.eglMakeCurrent(cdpy, cdraw, cread, cctx)
	__v := (Boolean)(__ret)
	return __v
}

// GetCurrentContext function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglGetCurrentContext.xhtml
func GetCurrentContext() Context {
	__ret := C.eglGetCurrentContext()
	__v := (Context)(__ret)
	return __v
}

// GetCurrentSurface function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglGetCurrentSurface.xhtml
func GetCurrentSurface(readdraw Int) Surface {
	creaddraw, _ := (C.EGLint)(readdraw), cgoAllocsUnknown
	__ret := C.eglGetCurrentSurface(creaddraw)
	__v := (Surface)(__ret)
	return __v
}

// GetCurrentDisplay function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglGetCurrentDisplay.xhtml
func GetCurrentDisplay() Display {
	__ret := C.eglGetCurrentDisplay()
	__v := (Display)(__ret)
	return __v
}

// QueryContext function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglQueryContext.xhtml
func QueryContext(dpy Display, ctx Context, attribute Int, value *Int) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	cctx, _ := (unsafe.Pointer)(ctx), cgoAllocsUnknown
	cattribute, _ := (C.EGLint)(attribute), cgoAllocsUnknown
	cvalue, _ := (*C.EGLint)(unsafe.Pointer(value)), cgoAllocsUnknown
	__ret := C.eglQueryContext(cdpy, cctx, cattribute, cvalue)
	__v := (Boolean)(__ret)
	return __v
}

// WaitGL function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglWaitGL.xhtml
func WaitGL() Boolean {
	__ret := C.eglWaitGL()
	__v := (Boolean)(__ret)
	return __v
}

// WaitNative function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglWaitNative.xhtml
func WaitNative(engine Int) Boolean {
	cengine, _ := (C.EGLint)(engine), cgoAllocsUnknown
	__ret := C.eglWaitNative(cengine)
	__v := (Boolean)(__ret)
	return __v
}

// SwapBuffers function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglSwapBuffers.xhtml
func SwapBuffers(dpy Display, surface Surface) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	csurface, _ := (unsafe.Pointer)(surface), cgoAllocsUnknown
	__ret := C.eglSwapBuffers(cdpy, csurface)
	__v := (Boolean)(__ret)
	return __v
}

// CopyBuffers function as declared in https://www.khronos.org/registry/egl/sdk/docs/man/html/eglCopyBuffers.xhtml
func CopyBuffers(dpy Display, surface Surface, target NativePixmapType) Boolean {
	cdpy, _ := (unsafe.Pointer)(dpy), cgoAllocsUnknown
	csurface, _ := (unsafe.Pointer)(surface), cgoAllocsUnknown
	ctarget, _ := *(*C.EGLNativePixmapType)(unsafe.Pointer(&target)), cgoAllocsUnknown
	__ret := C.eglCopyBuffers(cdpy, csurface, ctarget)
	__v := (Boolean)(__ret)
	return __v
}
