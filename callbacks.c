#include "callbacks.h"

int32_t go_plc_tag_create_ex(const char *attrib_str, 
    go_plc_tag_create_ex_callback callback_function, void *userdata, int timeout){
        return plc_tag_create_ex(attrib_str, callback_function, userdata, timeout);
}

int go_plc_tag_register_callback_ex(int32_t tag_id, go_plc_tag_register_ex_callback callback_function, void *userdata){
    return plc_tag_register_callback_ex(tag_id, callback_function, userdata);
}

int go_plc_tag_register_logger(go_plc_tag_register_logger_callback callback_function){
    return plc_tag_register_logger(callback_function);
}