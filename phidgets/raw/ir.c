#include <string.h>
#include "ir.h"

void onCodeResultFree(onCodeResult *r) {
  free(r->data);
  free(r);
}

void onLearnResultFree(onLearnResult *r) {
  free(r->data);
  free(r);
}

void onRawDataResultFree(onRawDataResult *r) {
  free(r->data);
  free(r);
}

onCodeResult * onCodeAwait(handler *h) {
  return (onCodeResult *)handlerAwait(h);
}

onLearnResult * onLearnAwait(handler *h) {
  return (onLearnResult *)handlerAwait(h);
}

onRawDataResult * onRawDataAwait(handler *h) {
  return (onRawDataResult *)handlerAwait(h);
}

int onCodeHandler(PhidgetIRHandle ir, void *ptr, unsigned char *data, int dataLength, int bitCount, int repeat) {
  handler * h = (handler *)ptr;

  onCodeResult * r = calloc(1, sizeof(onCodeResult));
  r->data = malloc(dataLength * sizeof(char));
  r->dataLength = dataLength;
  r->bitCount = bitCount;
  r->repeat = repeat;
  memcpy(r->data, data, dataLength * sizeof(char));
  handlerAppendResult(h, r);

  return 0;
}

int onLearnHandler(PhidgetIRHandle ir, void *ptr, unsigned char *data, int dataLength, PhidgetIR_CodeInfo codeInfo) {
  handler * h = (handler *)ptr;

  onLearnResult * r = calloc(1, sizeof(onLearnResult));
  r->data = malloc(dataLength * sizeof(char));
  r->dataLength = dataLength;
  r->codeInfo = codeInfo;
  memcpy(r->data, data, dataLength * sizeof(char));
  handlerAppendResult(h, r);

  return 0;
}

int onRawDataHandler(PhidgetIRHandle ir, void *ptr, int *data, int dataLength) {
  handler * h = (handler *)ptr;

  onRawDataResult * r = calloc(1, sizeof(onRawDataResult));
  r->data = malloc(dataLength * sizeof(int));
  r->dataLength = dataLength;
  memcpy(r->data, data, dataLength * sizeof(int));
  handlerAppendResult(h, r);

  return 0;
}

int setOnCodeHandler(PhidgetIRHandle ir, handler *h) {
  return PhidgetIR_setOnCodeHandler(ir, &onCodeHandler, h);
}

int setOnLearnHandler(PhidgetIRHandle ir, handler *h) {
  return PhidgetIR_setOnLearnHandler(ir, &onLearnHandler, h);
}

int setOnRawDataHandler(PhidgetIRHandle ir, handler *h) {
  return PhidgetIR_setOnRawDataHandler(ir, &onRawDataHandler, h);
}

void unsetOnCodeHandler(PhidgetIRHandle ir) {
  PhidgetIR_setOnCodeHandler(ir, NULL, NULL);
}

void unsetOnLearnHandler(PhidgetIRHandle ir) {
  PhidgetIR_setOnLearnHandler(ir, NULL, NULL);
}

void unsetOnRawDataHandler(PhidgetIRHandle ir) {
  PhidgetIR_setOnRawDataHandler(ir, NULL, NULL);
}
