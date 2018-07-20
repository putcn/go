#include "logging.h"

int _log(Phidget_LogLevel l, const char * message) {
  return PhidgetLog_log(l, message);
}
