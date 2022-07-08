#pragma once

#include <stdint.h>
#include "libplctag.h"

#ifdef __cplusplus
extern "C" {
#endif


#if  defined(_WIN32) || defined(WIN32) || defined(WIN64) || defined(_WIN64)
    #ifdef __cplusplus
        #define C_FUNC extern "C"
    #else
        #define C_FUNC
    #endif

    #ifdef LIBPLCTAGDLL_EXPORTS
        #define LIB_EXPORT __declspec(dllexport)
    #else
        #define LIB_EXPORT __declspec(dllimport)
    #endif
#else
    #ifdef LIBPLCTAGDLL_EXPORTS
        #define LIB_EXPORT __attribute__ ((visibility ("default")))
    #else
        #define LIB_EXPORT extern
    #endif
#endif

// int32_t go_plc_tag_create_ex(const char *attrib_str, void (*tag_callback_func)(int32_t tag_id, int event, int status, void *userdata), void *userdata, int timeout);
typedef void (*go_plc_tag_create_ex_callback)(int32_t tag_id, int event, int status, void *userdata);
LIB_EXPORT int32_t go_plc_tag_create_ex(const char *attrib_str, go_plc_tag_create_ex_callback callback_function, void *userdata, int timeout);

// int plc_tag_register_callback_ex(int32_t tag_id, void (*tag_callback_func)(int32_t tag_id, int event, int status, void *userdata), void *userdata);
typedef void (*go_plc_tag_register_ex_callback)(int32_t tag_id, int event, int status, void *userdata);
LIB_EXPORT int go_plc_tag_register_callback_ex(int32_t tag_id, go_plc_tag_register_ex_callback callback_function, void *userdata);

// int plc_tag_register_logger(void (*log_callback_func)(int32_t tag_id, int debug_level, const char *message));
typedef void (*go_plc_tag_register_logger_callback)(int32_t tag_id, int debug_level, const char *message);
LIB_EXPORT int go_plc_tag_register_logger(go_plc_tag_register_logger_callback callback_function);

#ifdef __cplusplus
}
#endif