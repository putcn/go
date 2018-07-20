#include <phidget22.h>
#include <stdlib.h>
#include "handler.h"

typedef struct onCodeResult {
  unsigned char *data;
  int dataLength;
  int bitCount;
  int repeat;
} onCodeResult;

typedef struct onLearnResult {
  unsigned char *data;
  int dataLength;
  PhidgetIR_CodeInfo codeInfo;
} onLearnResult;

typedef struct onRawDataResult {
  int *data;
  int dataLength;
} onRawDataResult;

void onCodeResultFree(onCodeResult *r);
void onLearnResultFree(onLearnResult *r);
void onRawDataResultFree(onRawDataResult *r);

onCodeResult * onCodeAwait(handler *h);
onLearnResult * onLearnAwait(handler *h);
onRawDataResult * onRawDataAwait(handler *h);

int setOnCodeHandler(PhidgetIRHandle ir, handler *h);
int setOnLearnHandler(PhidgetIRHandle ir, handler *h);
int setOnRawDataHandler(PhidgetIRHandle ir, handler *h);
void unsetOnCodeHandler(PhidgetIRHandle ir);
void unsetOnLearnHandler(PhidgetIRHandle ir);
void unsetOnRawDataHandler(PhidgetIRHandle ir);
