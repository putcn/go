#include <stddef.h>
#include <phidget22.h>
#include <stdlib.h>
#include "handler.h"

typedef enum {
  phidgetAttach = 1,
  phidgetConnect,
  phidgetDetach,
  phidgetDisconnect
} eventType;

typedef struct onErrorResult {
  Phidget_ErrorEventCode code;
  const char * string;
} onErrorResult;

void onErrorResultFree(onErrorResult *r);

onErrorResult * onErrorAwait(handler *h);
void onEventAwait(handler *h);

int setOnErrorHandler(PhidgetHandle p, handler *h);
int setOnEventHandler(PhidgetHandle p, handler *h, eventType t);
void unsetOnErrorHandler(PhidgetHandle p);
void unsetOnEventHandler(PhidgetHandle p, eventType t);
