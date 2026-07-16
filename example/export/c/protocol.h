// Auto-generated C protocol code
#ifndef PROTOCOL_H
#define PROTOCOL_H

#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "cjson/cJSON.h"

typedef struct ProtocolLogin ProtocolLogin;
typedef struct Account Account;
typedef struct ProtocolServerLogin ProtocolServerLogin;
typedef struct ProtocolServerState ProtocolServerState;
typedef struct ProtocolServerCommand ProtocolServerCommand;

struct ProtocolLogin {
    int8_t a;
    uint8_t b;
    int16_t c;
    uint16_t d;
    int32_t e;
    uint32_t f;
    int64_t g;
    uint64_t h;
    float i;
    double j;
    bool k;
    uint8_t l;
    char* m;
    int8_t*; int n_count n;
    char**; int o_count o;
    Account* q;
    Account*; int r_count r;
};

void ProtocolLogin_init(ProtocolLogin* self);
void ProtocolLogin_decode(ProtocolLogin* self, const uint8_t* data, size_t data_len);
uint8_t* ProtocolLogin_encode_binary(ProtocolLogin* self, size_t* out_len);
char* ProtocolLogin_encode_json(ProtocolLogin* self);
void ProtocolLogin_free(ProtocolLogin* self);

struct Account {
    char* nickname;
    char* password;
};

void Account_init(Account* self);
void Account_decode(Account* self, const uint8_t* data, size_t data_len);
uint8_t* Account_encode_binary(Account* self, size_t* out_len);
char* Account_encode_json(Account* self);
void Account_free(Account* self);

struct ProtocolServerLogin {
    char* id;
    char* token;
};

void ProtocolServerLogin_init(ProtocolServerLogin* self);
void ProtocolServerLogin_decode(ProtocolServerLogin* self, const uint8_t* data, size_t data_len);
uint8_t* ProtocolServerLogin_encode_binary(ProtocolServerLogin* self, size_t* out_len);
char* ProtocolServerLogin_encode_json(ProtocolServerLogin* self);
void ProtocolServerLogin_free(ProtocolServerLogin* self);

struct ProtocolServerState {
    uint16_t state;
    char* value;
};

void ProtocolServerState_init(ProtocolServerState* self);
void ProtocolServerState_decode(ProtocolServerState* self, const uint8_t* data, size_t data_len);
uint8_t* ProtocolServerState_encode_binary(ProtocolServerState* self, size_t* out_len);
char* ProtocolServerState_encode_json(ProtocolServerState* self);
void ProtocolServerState_free(ProtocolServerState* self);

struct ProtocolServerCommand {
    uint16_t command;
    char* value;
};

void ProtocolServerCommand_init(ProtocolServerCommand* self);
void ProtocolServerCommand_decode(ProtocolServerCommand* self, const uint8_t* data, size_t data_len);
uint8_t* ProtocolServerCommand_encode_binary(ProtocolServerCommand* self, size_t* out_len);
char* ProtocolServerCommand_encode_json(ProtocolServerCommand* self);
void ProtocolServerCommand_free(ProtocolServerCommand* self);

#endif // PROTOCOL_H


// Implementation
// Auto-generated C protocol implementation
#include "protocol.h"

void ProtocolLogin_init(ProtocolLogin* self) {
    if (self == NULL) return;
    memset(self, 0, sizeof(ProtocolLogin));
    self->m = NULL;
    self->n = NULL;
    self->n_count = 0;
    self->o = NULL;
    self->o_count = 0;
    self->q = (Account*)malloc(sizeof(Account));
    Account_init(self->q);
    self->r = NULL;
    self->r_count = 0;
}

void ProtocolLogin_decode(ProtocolLogin* self, const uint8_t* data, size_t data_len) {
    if (self == NULL || data == NULL || data_len == 0) return;
    
    // Check if JSON format
    if (data[0] == '{') {
        cJSON* json = cJSON_Parse((const char*)data);
        if (json != NULL) {
            cJSON* a_item = cJSON_GetObjectItemCaseSensitive(json, "a");
            if (cJSON_IsNumber(a_item)) {
                self->a = (int32_t)cJSON_GetNumberValue(a_item);
            }
            cJSON* b_item = cJSON_GetObjectItemCaseSensitive(json, "b");
            if (cJSON_IsNumber(b_item)) {
                self->b = (uint32_t)cJSON_GetNumberValue(b_item);
            }
            cJSON* c_item = cJSON_GetObjectItemCaseSensitive(json, "c");
            if (cJSON_IsNumber(c_item)) {
                self->c = (int32_t)cJSON_GetNumberValue(c_item);
            }
            cJSON* d_item = cJSON_GetObjectItemCaseSensitive(json, "d");
            if (cJSON_IsNumber(d_item)) {
                self->d = (uint32_t)cJSON_GetNumberValue(d_item);
            }
            cJSON* e_item = cJSON_GetObjectItemCaseSensitive(json, "e");
            if (cJSON_IsNumber(e_item)) {
                self->e = (int32_t)cJSON_GetNumberValue(e_item);
            }
            cJSON* f_item = cJSON_GetObjectItemCaseSensitive(json, "f");
            if (cJSON_IsNumber(f_item)) {
                self->f = (uint32_t)cJSON_GetNumberValue(f_item);
            }
            cJSON* g_item = cJSON_GetObjectItemCaseSensitive(json, "g");
            if (cJSON_IsNumber(g_item)) {
                self->g = (int64_t)cJSON_GetNumberValue(g_item);
            }
            cJSON* h_item = cJSON_GetObjectItemCaseSensitive(json, "h");
            if (cJSON_IsNumber(h_item)) {
                self->h = (int64_t)cJSON_GetNumberValue(h_item);
            }
            cJSON* i_item = cJSON_GetObjectItemCaseSensitive(json, "i");
            if (cJSON_IsNumber(i_item)) {
                self->i = (double)cJSON_GetNumberValue(i_item);
            }
            cJSON* j_item = cJSON_GetObjectItemCaseSensitive(json, "j");
            if (cJSON_IsNumber(j_item)) {
                self->j = (double)cJSON_GetNumberValue(j_item);
            }
            cJSON* k_item = cJSON_GetObjectItemCaseSensitive(json, "k");
            if (cJSON_IsBool(k_item)) {
                self->k = cJSON_IsTrue(k_item);
            }
            cJSON* l_item = cJSON_GetObjectItemCaseSensitive(json, "l");
            if (cJSON_IsNumber(l_item)) {
                self->l = (uint8_t)cJSON_GetNumberValue(l_item);
            }
            cJSON* m_item = cJSON_GetObjectItemCaseSensitive(json, "m");
            if (cJSON_IsString(m_item)) {
                const char* str = cJSON_GetStringValue(m_item);
                if (self->m) free(self->m);
                self->m = (char*)malloc(strlen(str) + 1);
                strcpy(self->m, str);
            }
            cJSON* n_item = cJSON_GetObjectItemCaseSensitive(json, "n");
            if (cJSON_IsArray(n_item)) {
                // Array decoding - simplified
            }
            cJSON* o_item = cJSON_GetObjectItemCaseSensitive(json, "o");
            if (cJSON_IsArray(o_item)) {
                // Array decoding - simplified
            }
            cJSON* q_item = cJSON_GetObjectItemCaseSensitive(json, "q");
            if (cJSON_IsObject(q_item)) {
                char* q_str = cJSON_Print(q_item);
                Account_decode(self->q, (uint8_t*)q_str, strlen(q_str));
                free(q_str);
            }
            cJSON* r_item = cJSON_GetObjectItemCaseSensitive(json, "r");
            if (cJSON_IsArray(r_item)) {
                // Array decoding - simplified
            }
            cJSON_Delete(json);
        }
        return;
    }
    
    // Binary decoding
    size_t pointer = 0;
    if (pointer + 1 <= data_len) {
        self->a = *(int8_t*)(data + pointer);
        pointer += 1;
    }
    if (pointer + 1 <= data_len) {
        self->b = *(uint8_t*)(data + pointer);
        pointer += 1;
    }
    if (pointer + 2 <= data_len) {
        self->c = *(int16_t*)(data + pointer);
        pointer += 2;
    }
    if (pointer + 2 <= data_len) {
        self->d = *(uint16_t*)(data + pointer);
        pointer += 2;
    }
    if (pointer + 4 <= data_len) {
        self->e = *(int32_t*)(data + pointer);
        pointer += 4;
    }
    if (pointer + 4 <= data_len) {
        self->f = *(uint32_t*)(data + pointer);
        pointer += 4;
    }
    if (pointer + 8 <= data_len) {
        self->g = *(int64_t*)(data + pointer);
        pointer += 8;
    }
    if (pointer + 8 <= data_len) {
        self->h = *(uint64_t*)(data + pointer);
        pointer += 8;
    }
    if (pointer + 4 <= data_len) {
        self->i = *(float*)(data + pointer);
        pointer += 4;
    }
    if (pointer + 8 <= data_len) {
        self->j = *(double*)(data + pointer);
        pointer += 8;
    }
    if (pointer + 1 <= data_len) {
        self->k = data[pointer] != 0;
        pointer += 1;
    }
    if (pointer + 1 <= data_len) {
        self->l = data[pointer];
        pointer += 1;
    }
    if (pointer + 4 <= data_len) {
        int32_t m_len = *(int32_t*)(data + pointer);
        pointer += 4;
        if (pointer + m_len <= data_len) {
            if (self->m) free(self->m);
            self->m = (char*)malloc(m_len + 1);
            memcpy(self->m, data + pointer, m_len);
            self->m[m_len] = '\0';
            pointer += m_len;
        }
    }
    if (pointer + 4 <= data_len) {
        int32_t n_count = *(int32_t*)(data + pointer);
        pointer += 4;
        if (self->n) free(self->n);
        self->n_count = n_count;
        self->n = (int8_t*)malloc(sizeof(int8_t) * n_count);
        for (int i = 0; i < n_count; i++) {
            if (pointer + 1 <= data_len) {
                n[i] = *(int8_t*)(data + pointer);
                pointer += 1;
            }
        }
    }
    if (pointer + 4 <= data_len) {
        int32_t o_count = *(int32_t*)(data + pointer);
        pointer += 4;
        if (self->o) free(self->o);
        self->o_count = o_count;
        self->o = (char**)malloc(sizeof(char*) * o_count);
        for (int i = 0; i < o_count; i++) {
            if (pointer + 4 <= data_len) {
                int32_t len = *(int32_t*)(data + pointer);
                pointer += 4;
                if (pointer + len <= data_len) {
                    self->o[i] = (char*)malloc(len + 1);
                    memcpy(self->o[i], data + pointer, len);
                    self->o[i][len] = '\0';
                    pointer += len;
                }
            }
        }
    }
    if (pointer + 4 <= data_len) {
        int32_t q_len = *(int32_t*)(data + pointer);
        pointer += 4;
        if (pointer + q_len <= data_len) {
            if (self->q == NULL) {
                self->q = (Account*)malloc(sizeof(Account));
                Account_init(self->q);
            }
            Account_decode(self->q, data + pointer, q_len);
            pointer += q_len;
        }
    }
    if (pointer + 4 <= data_len) {
        int32_t r_count = *(int32_t*)(data + pointer);
        pointer += 4;
        if (self->r) free(self->r);
        self->r_count = r_count;
        self->r = (Account*)malloc(sizeof(Account) * r_count);
        for (int i = 0; i < r_count; i++) {
            Account_init(&self->r[i]);
            if (pointer + 4 <= data_len) {
                int32_t len = *(int32_t*)(data + pointer);
                pointer += 4;
                if (pointer + len <= data_len) {
                    Account_decode(&self->r[i], data + pointer, len);
                    pointer += len;
                }
            }
        }
    }
}

uint8_t* ProtocolLogin_encode_binary(ProtocolLogin* self, size_t* out_len) {
    if (self == NULL) {
        if (out_len) *out_len = 0;
        return NULL;
    }
    
    size_t buffer_size = 1024;
    uint8_t* buffer = (uint8_t*)malloc(buffer_size);
    size_t offset = 0;
    
    // Encode a
    while (offset + 1 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->a, 1);
    offset += 1;
    // Encode b
    while (offset + 1 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->b, 1);
    offset += 1;
    // Encode c
    while (offset + 2 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->c, 2);
    offset += 2;
    // Encode d
    while (offset + 2 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->d, 2);
    offset += 2;
    // Encode e
    while (offset + 4 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->e, 4);
    offset += 4;
    // Encode f
    while (offset + 4 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->f, 4);
    offset += 4;
    // Encode g
    while (offset + 8 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->g, 8);
    offset += 8;
    // Encode h
    while (offset + 8 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->h, 8);
    offset += 8;
    // Encode i
    while (offset + 4 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->i, 4);
    offset += 4;
    // Encode j
    while (offset + 8 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->j, 8);
    offset += 8;
    // Encode k
    while (offset + 1 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->k, 1);
    offset += 1;
    // Encode l
    while (offset + 1 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->l, 1);
    offset += 1;
    // Encode m
    if (self->m) {
        size_t m_len = strlen(self->m);
        while (offset + 4 + m_len > buffer_size) {
            buffer_size *= 2;
            buffer = (uint8_t*)realloc(buffer, buffer_size);
        }
        int32_t len = (int32_t)m_len;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
        memcpy(buffer + offset, self->m, m_len);
        offset += m_len;
    } else {
        int32_t len = 0;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
    }
    // Encode n
    while (offset + 4 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    int32_t n_count = self->n_count;
    memcpy(buffer + offset, &n_count, 4);
    offset += 4;
    for (int i = 0; i < n_count; i++) {
        while (offset + 1 > buffer_size) {
            buffer_size *= 2;
            buffer = (uint8_t*)realloc(buffer, buffer_size);
        }
        memcpy(buffer + offset, &self->n[i], 1);
        offset += 1;
    }
    // Encode o
    while (offset + 4 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    int32_t o_count = self->o_count;
    memcpy(buffer + offset, &o_count, 4);
    offset += 4;
    for (int i = 0; i < o_count; i++) {
        if (self->o[i]) {
            size_t len = strlen(self->o[i]);
            while (offset + 4 + len > buffer_size) {
                buffer_size *= 2;
                buffer = (uint8_t*)realloc(buffer, buffer_size);
            }
            int32_t str_len = (int32_t)len;
            memcpy(buffer + offset, &str_len, 4);
            offset += 4;
            memcpy(buffer + offset, self->o[i], len);
            offset += len;
        } else {
            int32_t len = 0;
            memcpy(buffer + offset, &len, 4);
            offset += 4;
        }
    }
    // Encode q
    if (self->q) {
        size_t q_len = 0;
        uint8_t* q_data = Account_encode_binary(self->q, &q_len);
        if (q_data) {
            while (offset + 4 + q_len > buffer_size) {
                buffer_size *= 2;
                buffer = (uint8_t*)realloc(buffer, buffer_size);
            }
            int32_t len = (int32_t)q_len;
            memcpy(buffer + offset, &len, 4);
            offset += 4;
            memcpy(buffer + offset, q_data, q_len);
            offset += q_len;
            free(q_data);
        }
    } else {
        int32_t len = 0;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
    }
    // Encode r
    while (offset + 4 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    int32_t r_count = self->r_count;
    memcpy(buffer + offset, &r_count, 4);
    offset += 4;
    for (int i = 0; i < r_count; i++) {
        size_t item_len = 0;
        uint8_t* item_data = Account_encode_binary(&self->r[i], &item_len);
        if (item_data) {
            while (offset + 4 + item_len > buffer_size) {
                buffer_size *= 2;
                buffer = (uint8_t*)realloc(buffer, buffer_size);
            }
            int32_t len = (int32_t)item_len;
            memcpy(buffer + offset, &len, 4);
            offset += 4;
            memcpy(buffer + offset, item_data, item_len);
            offset += item_len;
            free(item_data);
        }
    }
    
    if (out_len) *out_len = offset;
    return buffer;
}

char* ProtocolLogin_encode_json(ProtocolLogin* self) {
    if (self == NULL) return NULL;
    
    cJSON* json = cJSON_CreateObject();
    cJSON_AddNumberToObject(json, "a", (double)self->a);
    cJSON_AddNumberToObject(json, "b", (double)self->b);
    cJSON_AddNumberToObject(json, "c", (double)self->c);
    cJSON_AddNumberToObject(json, "d", (double)self->d);
    cJSON_AddNumberToObject(json, "e", (double)self->e);
    cJSON_AddNumberToObject(json, "f", (double)self->f);
    cJSON_AddNumberToObject(json, "g", (double)self->g);
    cJSON_AddNumberToObject(json, "h", (double)self->h);
    cJSON_AddNumberToObject(json, "i", self->i);
    cJSON_AddNumberToObject(json, "j", self->j);
    cJSON_AddBoolToObject(json, "k", self->k);
    cJSON_AddNumberToObject(json, "l", (double)self->l);
    if (self->m) {
        cJSON_AddStringToObject(json, "m", self->m);
    } else {
        cJSON_AddStringToObject(json, "m", "");
    }
    cJSON* n_array = cJSON_CreateArray();
    for (int i = 0; i < self->n_count; i++) {
        cJSON_AddItemToArray(n_array, cJSON_CreateNumber((double)self->n[i]));
    }
    cJSON_AddItemToObject(json, "n", n_array);
    cJSON* o_array = cJSON_CreateArray();
    for (int i = 0; i < self->o_count; i++) {
        if (self->o[i]) {
            cJSON_AddItemToArray(o_array, cJSON_CreateString(self->o[i]));
        }
    }
    cJSON_AddItemToObject(json, "o", o_array);
    if (self->q) {
        char* q_json = Account_encode_json(self->q);
        if (q_json) {
            cJSON* q_obj = cJSON_Parse(q_json);
            if (q_obj) {
                cJSON_AddItemToObject(json, "q", q_obj);
            }
            free(q_json);
        }
    }
    cJSON* r_array = cJSON_CreateArray();
    for (int i = 0; i < self->r_count; i++) {
        cJSON_AddItemToArray(r_array, cJSON_CreateNumber((double)self->r[i]));
    }
    cJSON_AddItemToObject(json, "r", r_array);
    
    char* json_string = cJSON_Print(json);
    cJSON_Delete(json);
    return json_string;
}

void ProtocolLogin_free(ProtocolLogin* self) {
    if (self == NULL) return;
    if (self->m) free(self->m);
    if (self->n) free(self->n);
    if (self->o) {
        for (int i = 0; i < self->o_count; i++) {
            if (self->o[i]) free(self->o[i]);
        }
        free(self->o);
    }
    if (self->q) {
        Account_free(self->q);
        free(self->q);
    }
    if (self->r) {
        for (int i = 0; i < self->r_count; i++) {
            Account_free(&self->r[i]);
        }
        free(self->r);
    }
}

void Account_init(Account* self) {
    if (self == NULL) return;
    memset(self, 0, sizeof(Account));
    self->nickname = NULL;
    self->password = NULL;
}

void Account_decode(Account* self, const uint8_t* data, size_t data_len) {
    if (self == NULL || data == NULL || data_len == 0) return;
    
    // Check if JSON format
    if (data[0] == '{') {
        cJSON* json = cJSON_Parse((const char*)data);
        if (json != NULL) {
            cJSON* nickname_item = cJSON_GetObjectItemCaseSensitive(json, "nickname");
            if (cJSON_IsString(nickname_item)) {
                const char* str = cJSON_GetStringValue(nickname_item);
                if (self->nickname) free(self->nickname);
                self->nickname = (char*)malloc(strlen(str) + 1);
                strcpy(self->nickname, str);
            }
            cJSON* password_item = cJSON_GetObjectItemCaseSensitive(json, "password");
            if (cJSON_IsString(password_item)) {
                const char* str = cJSON_GetStringValue(password_item);
                if (self->password) free(self->password);
                self->password = (char*)malloc(strlen(str) + 1);
                strcpy(self->password, str);
            }
            cJSON_Delete(json);
        }
        return;
    }
    
    // Binary decoding
    size_t pointer = 0;
    if (pointer + 4 <= data_len) {
        int32_t nickname_len = *(int32_t*)(data + pointer);
        pointer += 4;
        if (pointer + nickname_len <= data_len) {
            if (self->nickname) free(self->nickname);
            self->nickname = (char*)malloc(nickname_len + 1);
            memcpy(self->nickname, data + pointer, nickname_len);
            self->nickname[nickname_len] = '\0';
            pointer += nickname_len;
        }
    }
    if (pointer + 4 <= data_len) {
        int32_t password_len = *(int32_t*)(data + pointer);
        pointer += 4;
        if (pointer + password_len <= data_len) {
            if (self->password) free(self->password);
            self->password = (char*)malloc(password_len + 1);
            memcpy(self->password, data + pointer, password_len);
            self->password[password_len] = '\0';
            pointer += password_len;
        }
    }
}

uint8_t* Account_encode_binary(Account* self, size_t* out_len) {
    if (self == NULL) {
        if (out_len) *out_len = 0;
        return NULL;
    }
    
    size_t buffer_size = 1024;
    uint8_t* buffer = (uint8_t*)malloc(buffer_size);
    size_t offset = 0;
    
    // Encode nickname
    if (self->nickname) {
        size_t nickname_len = strlen(self->nickname);
        while (offset + 4 + nickname_len > buffer_size) {
            buffer_size *= 2;
            buffer = (uint8_t*)realloc(buffer, buffer_size);
        }
        int32_t len = (int32_t)nickname_len;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
        memcpy(buffer + offset, self->nickname, nickname_len);
        offset += nickname_len;
    } else {
        int32_t len = 0;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
    }
    // Encode password
    if (self->password) {
        size_t password_len = strlen(self->password);
        while (offset + 4 + password_len > buffer_size) {
            buffer_size *= 2;
            buffer = (uint8_t*)realloc(buffer, buffer_size);
        }
        int32_t len = (int32_t)password_len;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
        memcpy(buffer + offset, self->password, password_len);
        offset += password_len;
    } else {
        int32_t len = 0;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
    }
    
    if (out_len) *out_len = offset;
    return buffer;
}

char* Account_encode_json(Account* self) {
    if (self == NULL) return NULL;
    
    cJSON* json = cJSON_CreateObject();
    if (self->nickname) {
        cJSON_AddStringToObject(json, "nickname", self->nickname);
    } else {
        cJSON_AddStringToObject(json, "nickname", "");
    }
    if (self->password) {
        cJSON_AddStringToObject(json, "password", self->password);
    } else {
        cJSON_AddStringToObject(json, "password", "");
    }
    
    char* json_string = cJSON_Print(json);
    cJSON_Delete(json);
    return json_string;
}

void Account_free(Account* self) {
    if (self == NULL) return;
    if (self->nickname) free(self->nickname);
    if (self->password) free(self->password);
}

void ProtocolServerLogin_init(ProtocolServerLogin* self) {
    if (self == NULL) return;
    memset(self, 0, sizeof(ProtocolServerLogin));
    self->id = NULL;
    self->token = NULL;
}

void ProtocolServerLogin_decode(ProtocolServerLogin* self, const uint8_t* data, size_t data_len) {
    if (self == NULL || data == NULL || data_len == 0) return;
    
    // Check if JSON format
    if (data[0] == '{') {
        cJSON* json = cJSON_Parse((const char*)data);
        if (json != NULL) {
            cJSON* id_item = cJSON_GetObjectItemCaseSensitive(json, "id");
            if (cJSON_IsString(id_item)) {
                const char* str = cJSON_GetStringValue(id_item);
                if (self->id) free(self->id);
                self->id = (char*)malloc(strlen(str) + 1);
                strcpy(self->id, str);
            }
            cJSON* token_item = cJSON_GetObjectItemCaseSensitive(json, "token");
            if (cJSON_IsString(token_item)) {
                const char* str = cJSON_GetStringValue(token_item);
                if (self->token) free(self->token);
                self->token = (char*)malloc(strlen(str) + 1);
                strcpy(self->token, str);
            }
            cJSON_Delete(json);
        }
        return;
    }
    
    // Binary decoding
    size_t pointer = 0;
    if (pointer + 4 <= data_len) {
        int32_t id_len = *(int32_t*)(data + pointer);
        pointer += 4;
        if (pointer + id_len <= data_len) {
            if (self->id) free(self->id);
            self->id = (char*)malloc(id_len + 1);
            memcpy(self->id, data + pointer, id_len);
            self->id[id_len] = '\0';
            pointer += id_len;
        }
    }
    if (pointer + 4 <= data_len) {
        int32_t token_len = *(int32_t*)(data + pointer);
        pointer += 4;
        if (pointer + token_len <= data_len) {
            if (self->token) free(self->token);
            self->token = (char*)malloc(token_len + 1);
            memcpy(self->token, data + pointer, token_len);
            self->token[token_len] = '\0';
            pointer += token_len;
        }
    }
}

uint8_t* ProtocolServerLogin_encode_binary(ProtocolServerLogin* self, size_t* out_len) {
    if (self == NULL) {
        if (out_len) *out_len = 0;
        return NULL;
    }
    
    size_t buffer_size = 1024;
    uint8_t* buffer = (uint8_t*)malloc(buffer_size);
    size_t offset = 0;
    
    // Encode id
    if (self->id) {
        size_t id_len = strlen(self->id);
        while (offset + 4 + id_len > buffer_size) {
            buffer_size *= 2;
            buffer = (uint8_t*)realloc(buffer, buffer_size);
        }
        int32_t len = (int32_t)id_len;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
        memcpy(buffer + offset, self->id, id_len);
        offset += id_len;
    } else {
        int32_t len = 0;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
    }
    // Encode token
    if (self->token) {
        size_t token_len = strlen(self->token);
        while (offset + 4 + token_len > buffer_size) {
            buffer_size *= 2;
            buffer = (uint8_t*)realloc(buffer, buffer_size);
        }
        int32_t len = (int32_t)token_len;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
        memcpy(buffer + offset, self->token, token_len);
        offset += token_len;
    } else {
        int32_t len = 0;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
    }
    
    if (out_len) *out_len = offset;
    return buffer;
}

char* ProtocolServerLogin_encode_json(ProtocolServerLogin* self) {
    if (self == NULL) return NULL;
    
    cJSON* json = cJSON_CreateObject();
    if (self->id) {
        cJSON_AddStringToObject(json, "id", self->id);
    } else {
        cJSON_AddStringToObject(json, "id", "");
    }
    if (self->token) {
        cJSON_AddStringToObject(json, "token", self->token);
    } else {
        cJSON_AddStringToObject(json, "token", "");
    }
    
    char* json_string = cJSON_Print(json);
    cJSON_Delete(json);
    return json_string;
}

void ProtocolServerLogin_free(ProtocolServerLogin* self) {
    if (self == NULL) return;
    if (self->id) free(self->id);
    if (self->token) free(self->token);
}

void ProtocolServerState_init(ProtocolServerState* self) {
    if (self == NULL) return;
    memset(self, 0, sizeof(ProtocolServerState));
    self->value = NULL;
}

void ProtocolServerState_decode(ProtocolServerState* self, const uint8_t* data, size_t data_len) {
    if (self == NULL || data == NULL || data_len == 0) return;
    
    // Check if JSON format
    if (data[0] == '{') {
        cJSON* json = cJSON_Parse((const char*)data);
        if (json != NULL) {
            cJSON* state_item = cJSON_GetObjectItemCaseSensitive(json, "state");
            if (cJSON_IsNumber(state_item)) {
                self->state = (uint32_t)cJSON_GetNumberValue(state_item);
            }
            cJSON* value_item = cJSON_GetObjectItemCaseSensitive(json, "value");
            if (cJSON_IsString(value_item)) {
                const char* str = cJSON_GetStringValue(value_item);
                if (self->value) free(self->value);
                self->value = (char*)malloc(strlen(str) + 1);
                strcpy(self->value, str);
            }
            cJSON_Delete(json);
        }
        return;
    }
    
    // Binary decoding
    size_t pointer = 0;
    if (pointer + 2 <= data_len) {
        self->state = *(uint16_t*)(data + pointer);
        pointer += 2;
    }
    if (pointer + 4 <= data_len) {
        int32_t value_len = *(int32_t*)(data + pointer);
        pointer += 4;
        if (pointer + value_len <= data_len) {
            if (self->value) free(self->value);
            self->value = (char*)malloc(value_len + 1);
            memcpy(self->value, data + pointer, value_len);
            self->value[value_len] = '\0';
            pointer += value_len;
        }
    }
}

uint8_t* ProtocolServerState_encode_binary(ProtocolServerState* self, size_t* out_len) {
    if (self == NULL) {
        if (out_len) *out_len = 0;
        return NULL;
    }
    
    size_t buffer_size = 1024;
    uint8_t* buffer = (uint8_t*)malloc(buffer_size);
    size_t offset = 0;
    
    // Encode state
    while (offset + 2 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->state, 2);
    offset += 2;
    // Encode value
    if (self->value) {
        size_t value_len = strlen(self->value);
        while (offset + 4 + value_len > buffer_size) {
            buffer_size *= 2;
            buffer = (uint8_t*)realloc(buffer, buffer_size);
        }
        int32_t len = (int32_t)value_len;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
        memcpy(buffer + offset, self->value, value_len);
        offset += value_len;
    } else {
        int32_t len = 0;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
    }
    
    if (out_len) *out_len = offset;
    return buffer;
}

char* ProtocolServerState_encode_json(ProtocolServerState* self) {
    if (self == NULL) return NULL;
    
    cJSON* json = cJSON_CreateObject();
    cJSON_AddNumberToObject(json, "state", (double)self->state);
    if (self->value) {
        cJSON_AddStringToObject(json, "value", self->value);
    } else {
        cJSON_AddStringToObject(json, "value", "");
    }
    
    char* json_string = cJSON_Print(json);
    cJSON_Delete(json);
    return json_string;
}

void ProtocolServerState_free(ProtocolServerState* self) {
    if (self == NULL) return;
    if (self->value) free(self->value);
}

void ProtocolServerCommand_init(ProtocolServerCommand* self) {
    if (self == NULL) return;
    memset(self, 0, sizeof(ProtocolServerCommand));
    self->value = NULL;
}

void ProtocolServerCommand_decode(ProtocolServerCommand* self, const uint8_t* data, size_t data_len) {
    if (self == NULL || data == NULL || data_len == 0) return;
    
    // Check if JSON format
    if (data[0] == '{') {
        cJSON* json = cJSON_Parse((const char*)data);
        if (json != NULL) {
            cJSON* command_item = cJSON_GetObjectItemCaseSensitive(json, "command");
            if (cJSON_IsNumber(command_item)) {
                self->command = (uint32_t)cJSON_GetNumberValue(command_item);
            }
            cJSON* value_item = cJSON_GetObjectItemCaseSensitive(json, "value");
            if (cJSON_IsString(value_item)) {
                const char* str = cJSON_GetStringValue(value_item);
                if (self->value) free(self->value);
                self->value = (char*)malloc(strlen(str) + 1);
                strcpy(self->value, str);
            }
            cJSON_Delete(json);
        }
        return;
    }
    
    // Binary decoding
    size_t pointer = 0;
    if (pointer + 2 <= data_len) {
        self->command = *(uint16_t*)(data + pointer);
        pointer += 2;
    }
    if (pointer + 4 <= data_len) {
        int32_t value_len = *(int32_t*)(data + pointer);
        pointer += 4;
        if (pointer + value_len <= data_len) {
            if (self->value) free(self->value);
            self->value = (char*)malloc(value_len + 1);
            memcpy(self->value, data + pointer, value_len);
            self->value[value_len] = '\0';
            pointer += value_len;
        }
    }
}

uint8_t* ProtocolServerCommand_encode_binary(ProtocolServerCommand* self, size_t* out_len) {
    if (self == NULL) {
        if (out_len) *out_len = 0;
        return NULL;
    }
    
    size_t buffer_size = 1024;
    uint8_t* buffer = (uint8_t*)malloc(buffer_size);
    size_t offset = 0;
    
    // Encode command
    while (offset + 2 > buffer_size) {
        buffer_size *= 2;
        buffer = (uint8_t*)realloc(buffer, buffer_size);
    }
    memcpy(buffer + offset, &self->command, 2);
    offset += 2;
    // Encode value
    if (self->value) {
        size_t value_len = strlen(self->value);
        while (offset + 4 + value_len > buffer_size) {
            buffer_size *= 2;
            buffer = (uint8_t*)realloc(buffer, buffer_size);
        }
        int32_t len = (int32_t)value_len;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
        memcpy(buffer + offset, self->value, value_len);
        offset += value_len;
    } else {
        int32_t len = 0;
        memcpy(buffer + offset, &len, 4);
        offset += 4;
    }
    
    if (out_len) *out_len = offset;
    return buffer;
}

char* ProtocolServerCommand_encode_json(ProtocolServerCommand* self) {
    if (self == NULL) return NULL;
    
    cJSON* json = cJSON_CreateObject();
    cJSON_AddNumberToObject(json, "command", (double)self->command);
    if (self->value) {
        cJSON_AddStringToObject(json, "value", self->value);
    } else {
        cJSON_AddStringToObject(json, "value", "");
    }
    
    char* json_string = cJSON_Print(json);
    cJSON_Delete(json);
    return json_string;
}

void ProtocolServerCommand_free(ProtocolServerCommand* self) {
    if (self == NULL) return;
    if (self->value) free(self->value);
}

