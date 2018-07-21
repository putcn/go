#include "phidget.h"

void onErrorResultFree(onErrorResult * r) {
  free(r);
}

onErrorResult * onErrorAwait(handler *h) {
  return (onErrorResult *)handlerAwait(h);
}

void onEventAwait(handler *h) {
  handlerAwait(h);
}

int onErrorHandler(PhidgetHandle p, void *ptr, Phidget_ErrorEventCode code, const char *string) {
  handler * h = (handler *)ptr;

  onErrorResult * r = calloc(1, sizeof(onErrorResult));
  r->code = code;
  r->string = string;
  handlerAppendResult(h, r);

  return 0;
}

int onEventHandler(PhidgetHandle p, void *ptr) {
  handler * h = (handler *)ptr;
  handlerAppendResult(h, NULL);
  return 0;
}

int setOnErrorHandler(PhidgetHandle p, handler *h) {
  return Phidget_setOnErrorHandler(p, &onErrorHandler, h);
}

int setOnEventHandler(PhidgetHandle p, handler *h, eventType t) {
  switch (t) {
    case phidgetAttach:
      return Phidget_setOnAttachHandler(p, &onEventHandler, h);
    case phidgetConnect:
      return 1; //CPhidget_set_OnServerConnect_Handler(p, &onEventHandler, h);
    case phidgetDetach:
      return Phidget_setOnDetachHandler(p, &onEventHandler, h);
    case phidgetDisconnect:
      return 1; //CPhidget_set_OnServerDisconnect_Handler(p, &onEventHandler, h);
  }
}

void unsetOnErrorHandler(PhidgetHandle p) {
  Phidget_setOnErrorHandler(p, NULL, NULL);
}

void unsetOnEventHandler(PhidgetHandle p, eventType t) {
  switch (t) {
    case phidgetAttach:
      Phidget_setOnAttachHandler(p, NULL, NULL);
    case phidgetConnect:
      ;//CPhidget_set_OnServerConnect_Handler(p, NULL, NULL);
    case phidgetDetach:
      Phidget_setOnDetachHandler(p, NULL, NULL);
    case phidgetDisconnect:
      ;//CPhidget_set_OnServerDisconnect_Handler(p, NULL, NULL);
  }
}
