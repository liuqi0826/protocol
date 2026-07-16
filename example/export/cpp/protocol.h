// Auto-generated C++ protocol code
#ifndef PROTOCOL_H
#define PROTOCOL_H

#include <cstdint>
#include <string>
#include <vector>
#include <memory>
#include <cstring>
#include <iostream>
#include "nlohmann/json.hpp"

using json = nlohmann::json;

class ProtocolLogin;
class Account;
class ProtocolServerLogin;
class ProtocolServerState;
class ProtocolServerCommand;

class ProtocolLogin {
public:
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
    std::string m;
    std::vector<int8_t> n;
    std::vector<std::string> o;
    Account q;
    std::vector<Account> r;

    ProtocolLogin() {
        // All members are default initialized
    }

    void decode(const uint8_t* data, size_t data_len) {
        if (data == nullptr || data_len == 0) return;
        
        // Check if JSON format
        if (data[0] == '{') {
            try {
                std::string json_str(reinterpret_cast<const char*>(data), data_len);
                json j = json::parse(json_str);
                if (j.contains("a")) {
                    this->a = j["a"].get<int64_t>();
                }
                if (j.contains("b")) {
                    this->b = j["b"].get<uint64_t>();
                }
                if (j.contains("c")) {
                    this->c = j["c"].get<int64_t>();
                }
                if (j.contains("d")) {
                    this->d = j["d"].get<uint64_t>();
                }
                if (j.contains("e")) {
                    this->e = j["e"].get<int64_t>();
                }
                if (j.contains("f")) {
                    this->f = j["f"].get<uint64_t>();
                }
                if (j.contains("g")) {
                    this->g = j["g"].get<int64_t>();
                }
                if (j.contains("h")) {
                    this->h = j["h"].get<uint64_t>();
                }
                if (j.contains("i")) {
                    this->i = j["i"].get<float>();
                }
                if (j.contains("j")) {
                    this->j = j["j"].get<double>();
                }
                if (j.contains("k")) {
                    this->k = j["k"].get<bool>();
                }
                if (j.contains("l")) {
                    this->l = j["l"].get<uint8_t>();
                }
                if (j.contains("m")) {
                    this->m = j["m"].get<std::string>();
                }
                if (j.contains("n")) {
                    auto arr = j["n"];
                    this->n.clear();
                    for (auto& item : arr) {
                        this->n.push_back(item.get<int8_t>());
                    }
                }
                if (j.contains("o")) {
                    this->o = j["o"].get<std::vector<std::string>>();
                }
                if (j.contains("q")) {
                    std::string q_str = j["q"].dump();
                    this->q.decode(reinterpret_cast<const uint8_t*>(q_str.c_str()), q_str.length());
                }
                if (j.contains("r")) {
                    auto arr = j["r"];
                    this->r.clear();
                    for (auto& item : arr) {
                        Account obj;
                        std::string item_str = item.dump();
                        obj.decode(reinterpret_cast<const uint8_t*>(item_str.c_str()), item_str.length());
                        this->r.push_back(obj);
                    }
                }
            } catch (const std::exception& e) {
                std::cerr << "JSON parse error: " << e.what() << std::endl;
            }
            return;
        }
        
        // Binary decoding
        size_t pointer = 0;
        if (pointer + 1 <= data_len) {
            a = *reinterpret_cast<const int8_t*>(data + pointer);
            pointer += 1;
        }
        if (pointer + 1 <= data_len) {
            b = data[pointer];
            pointer += 1;
        }
        if (pointer + 2 <= data_len) {
            c = *reinterpret_cast<const int16_t*>(data + pointer);
            pointer += 2;
        }
        if (pointer + 2 <= data_len) {
            d = *reinterpret_cast<const uint16_t*>(data + pointer);
            pointer += 2;
        }
        if (pointer + 4 <= data_len) {
            e = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
        }
        if (pointer + 4 <= data_len) {
            f = *reinterpret_cast<const uint32_t*>(data + pointer);
            pointer += 4;
        }
        if (pointer + 8 <= data_len) {
            g = *reinterpret_cast<const int64_t*>(data + pointer);
            pointer += 8;
        }
        if (pointer + 8 <= data_len) {
            h = *reinterpret_cast<const uint64_t*>(data + pointer);
            pointer += 8;
        }
        if (pointer + 4 <= data_len) {
            i = *reinterpret_cast<const float*>(data + pointer);
            pointer += 4;
        }
        if (pointer + 8 <= data_len) {
            j = *reinterpret_cast<const double*>(data + pointer);
            pointer += 8;
        }
        if (pointer + 1 <= data_len) {
            k = data[pointer] != 0;
            pointer += 1;
        }
        if (pointer + 1 <= data_len) {
            l = data[pointer];
            pointer += 1;
        }
        if (pointer + 4 <= data_len) {
            int32_t m_len = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            if (pointer + m_len <= data_len) {
                m = std::string(reinterpret_cast<const char*>(data + pointer), m_len);
                pointer += m_len;
            }
        }
        if (pointer + 4 <= data_len) {
            int32_t n_count = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            n.clear();
            n.reserve(n_count);
            for (int i = 0; i < n_count; i++) {
                if (pointer + 1 <= data_len) {
                    n.push_back(*reinterpret_cast<const int8_t*>(data + pointer));
                    pointer += 1;
                }
            }
        }
        if (pointer + 4 <= data_len) {
            int32_t o_count = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            o.clear();
            o.reserve(o_count);
            for (int i = 0; i < o_count; i++) {
                if (pointer + 4 <= data_len) {
                    int32_t len = *reinterpret_cast<const int32_t*>(data + pointer);
                    pointer += 4;
                    if (pointer + len <= data_len) {
                        o.push_back(std::string(reinterpret_cast<const char*>(data + pointer), len));
                        pointer += len;
                    }
                }
            }
        }
        if (pointer + 4 <= data_len) {
            int32_t q_len = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            if (pointer + q_len <= data_len) {
                q.decode(data + pointer, q_len);
                pointer += q_len;
            }
        }
        if (pointer + 4 <= data_len) {
            int32_t r_count = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            r.clear();
            r.reserve(r_count);
            for (int i = 0; i < r_count; i++) {
                if (pointer + 4 <= data_len) {
                    int32_t len = *reinterpret_cast<const int32_t*>(data + pointer);
                    pointer += 4;
                    if (pointer + len <= data_len) {
                        Account obj;
                        obj.decode(data + pointer, len);
                        r.push_back(obj);
                        pointer += len;
                    }
                }
            }
        }
    }

    std::vector<uint8_t> encodeBinary() const {
        std::vector<uint8_t> buffer;
        // Encode a
        {
            const int8_t* ptr = &a;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 1);
        }
        // Encode b
        {
            const uint8_t* ptr = &b;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 1);
        }
        // Encode c
        {
            const int16_t* ptr = &c;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 2);
        }
        // Encode d
        {
            const uint16_t* ptr = &d;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 2);
        }
        // Encode e
        {
            const int32_t* ptr = &e;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 4);
        }
        // Encode f
        {
            const uint32_t* ptr = &f;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 4);
        }
        // Encode g
        {
            const int64_t* ptr = &g;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 8);
        }
        // Encode h
        {
            const uint64_t* ptr = &h;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 8);
        }
        // Encode i
        {
            const float* ptr = &i;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 4);
        }
        // Encode j
        {
            const double* ptr = &j;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 8);
        }
        // Encode k
        {
            const bool* ptr = &k;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 1);
        }
        // Encode l
        {
            const uint8_t* ptr = &l;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 1);
        }
        // Encode m
        {
            int32_t len = static_cast<int32_t>(m.length());
            const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);
            buffer.insert(buffer.end(), len_bytes, len_bytes + 4);
            const uint8_t* str_bytes = reinterpret_cast<const uint8_t*>(m.c_str());
            buffer.insert(buffer.end(), str_bytes, str_bytes + len);
        }
        // Encode n
        {
            int32_t count = static_cast<int32_t>(n.size());
            const uint8_t* count_bytes = reinterpret_cast<const uint8_t*>(&count);
            buffer.insert(buffer.end(), count_bytes, count_bytes + 4);
            for (const auto& item : n) {
                const int8_t* ptr = &item;
                const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
                buffer.insert(buffer.end(), bytes, bytes + 1);
            }
        }
        // Encode o
        {
            int32_t count = static_cast<int32_t>(o.size());
            const uint8_t* count_bytes = reinterpret_cast<const uint8_t*>(&count);
            buffer.insert(buffer.end(), count_bytes, count_bytes + 4);
            for (const auto& item : o) {
                int32_t len = static_cast<int32_t>(item.length());
                const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);
                buffer.insert(buffer.end(), len_bytes, len_bytes + 4);
                const uint8_t* str_bytes = reinterpret_cast<const uint8_t*>(item.c_str());
                buffer.insert(buffer.end(), str_bytes, str_bytes + len);
            }
        }
        // Encode q
        {
            std::vector<uint8_t> q_data = q.encodeBinary();
            int32_t len = static_cast<int32_t>(q_data.size());
            const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);
            buffer.insert(buffer.end(), len_bytes, len_bytes + 4);
            buffer.insert(buffer.end(), q_data.begin(), q_data.end());
        }
        // Encode r
        {
            int32_t count = static_cast<int32_t>(r.size());
            const uint8_t* count_bytes = reinterpret_cast<const uint8_t*>(&count);
            buffer.insert(buffer.end(), count_bytes, count_bytes + 4);
            for (const auto& item : r) {
                std::vector<uint8_t> item_data = item.encodeBinary();
                int32_t len = static_cast<int32_t>(item_data.size());
                const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);
                buffer.insert(buffer.end(), len_bytes, len_bytes + 4);
                buffer.insert(buffer.end(), item_data.begin(), item_data.end());
            }
        }
        return buffer;
    }

    std::string encodeJson() const {
        json j;
        j["a"] = static_cast<int64_t>(a);
        j["b"] = static_cast<uint64_t>(b);
        j["c"] = static_cast<int64_t>(c);
        j["d"] = static_cast<uint64_t>(d);
        j["e"] = static_cast<int64_t>(e);
        j["f"] = static_cast<uint64_t>(f);
        j["g"] = static_cast<int64_t>(g);
        j["h"] = static_cast<uint64_t>(h);
        j["i"] = i;
        j["j"] = j;
        j["k"] = k;
        j["l"] = static_cast<uint8_t>(l);
        j["m"] = m;
        j["n"] = n;
        j["o"] = o;
        std::string q_json = q.encodeJson();
        j["q"] = json::parse(q_json);
        j["r"] = json::array();
        for (const auto& item : r) {
            std::string item_json = item.encodeJson();
            j["r"].push_back(json::parse(item_json));
        }
        return j.dump();
    }
};

class Account {
public:
    std::string nickname;
    std::string password;

    Account() {
        // All members are default initialized
    }

    void decode(const uint8_t* data, size_t data_len) {
        if (data == nullptr || data_len == 0) return;
        
        // Check if JSON format
        if (data[0] == '{') {
            try {
                std::string json_str(reinterpret_cast<const char*>(data), data_len);
                json j = json::parse(json_str);
                if (j.contains("nickname")) {
                    this->nickname = j["nickname"].get<std::string>();
                }
                if (j.contains("password")) {
                    this->password = j["password"].get<std::string>();
                }
            } catch (const std::exception& e) {
                std::cerr << "JSON parse error: " << e.what() << std::endl;
            }
            return;
        }
        
        // Binary decoding
        size_t pointer = 0;
        if (pointer + 4 <= data_len) {
            int32_t nickname_len = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            if (pointer + nickname_len <= data_len) {
                nickname = std::string(reinterpret_cast<const char*>(data + pointer), nickname_len);
                pointer += nickname_len;
            }
        }
        if (pointer + 4 <= data_len) {
            int32_t password_len = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            if (pointer + password_len <= data_len) {
                password = std::string(reinterpret_cast<const char*>(data + pointer), password_len);
                pointer += password_len;
            }
        }
    }

    std::vector<uint8_t> encodeBinary() const {
        std::vector<uint8_t> buffer;
        // Encode nickname
        {
            int32_t len = static_cast<int32_t>(nickname.length());
            const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);
            buffer.insert(buffer.end(), len_bytes, len_bytes + 4);
            const uint8_t* str_bytes = reinterpret_cast<const uint8_t*>(nickname.c_str());
            buffer.insert(buffer.end(), str_bytes, str_bytes + len);
        }
        // Encode password
        {
            int32_t len = static_cast<int32_t>(password.length());
            const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);
            buffer.insert(buffer.end(), len_bytes, len_bytes + 4);
            const uint8_t* str_bytes = reinterpret_cast<const uint8_t*>(password.c_str());
            buffer.insert(buffer.end(), str_bytes, str_bytes + len);
        }
        return buffer;
    }

    std::string encodeJson() const {
        json j;
        j["nickname"] = nickname;
        j["password"] = password;
        return j.dump();
    }
};

class ProtocolServerLogin {
public:
    std::string id;
    std::string token;

    ProtocolServerLogin() {
        // All members are default initialized
    }

    void decode(const uint8_t* data, size_t data_len) {
        if (data == nullptr || data_len == 0) return;
        
        // Check if JSON format
        if (data[0] == '{') {
            try {
                std::string json_str(reinterpret_cast<const char*>(data), data_len);
                json j = json::parse(json_str);
                if (j.contains("id")) {
                    this->id = j["id"].get<std::string>();
                }
                if (j.contains("token")) {
                    this->token = j["token"].get<std::string>();
                }
            } catch (const std::exception& e) {
                std::cerr << "JSON parse error: " << e.what() << std::endl;
            }
            return;
        }
        
        // Binary decoding
        size_t pointer = 0;
        if (pointer + 4 <= data_len) {
            int32_t id_len = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            if (pointer + id_len <= data_len) {
                id = std::string(reinterpret_cast<const char*>(data + pointer), id_len);
                pointer += id_len;
            }
        }
        if (pointer + 4 <= data_len) {
            int32_t token_len = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            if (pointer + token_len <= data_len) {
                token = std::string(reinterpret_cast<const char*>(data + pointer), token_len);
                pointer += token_len;
            }
        }
    }

    std::vector<uint8_t> encodeBinary() const {
        std::vector<uint8_t> buffer;
        // Encode id
        {
            int32_t len = static_cast<int32_t>(id.length());
            const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);
            buffer.insert(buffer.end(), len_bytes, len_bytes + 4);
            const uint8_t* str_bytes = reinterpret_cast<const uint8_t*>(id.c_str());
            buffer.insert(buffer.end(), str_bytes, str_bytes + len);
        }
        // Encode token
        {
            int32_t len = static_cast<int32_t>(token.length());
            const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);
            buffer.insert(buffer.end(), len_bytes, len_bytes + 4);
            const uint8_t* str_bytes = reinterpret_cast<const uint8_t*>(token.c_str());
            buffer.insert(buffer.end(), str_bytes, str_bytes + len);
        }
        return buffer;
    }

    std::string encodeJson() const {
        json j;
        j["id"] = id;
        j["token"] = token;
        return j.dump();
    }
};

class ProtocolServerState {
public:
    uint16_t state;
    std::string value;

    ProtocolServerState() {
        // All members are default initialized
    }

    void decode(const uint8_t* data, size_t data_len) {
        if (data == nullptr || data_len == 0) return;
        
        // Check if JSON format
        if (data[0] == '{') {
            try {
                std::string json_str(reinterpret_cast<const char*>(data), data_len);
                json j = json::parse(json_str);
                if (j.contains("state")) {
                    this->state = j["state"].get<uint64_t>();
                }
                if (j.contains("value")) {
                    this->value = j["value"].get<std::string>();
                }
            } catch (const std::exception& e) {
                std::cerr << "JSON parse error: " << e.what() << std::endl;
            }
            return;
        }
        
        // Binary decoding
        size_t pointer = 0;
        if (pointer + 2 <= data_len) {
            state = *reinterpret_cast<const uint16_t*>(data + pointer);
            pointer += 2;
        }
        if (pointer + 4 <= data_len) {
            int32_t value_len = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            if (pointer + value_len <= data_len) {
                value = std::string(reinterpret_cast<const char*>(data + pointer), value_len);
                pointer += value_len;
            }
        }
    }

    std::vector<uint8_t> encodeBinary() const {
        std::vector<uint8_t> buffer;
        // Encode state
        {
            const uint16_t* ptr = &state;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 2);
        }
        // Encode value
        {
            int32_t len = static_cast<int32_t>(value.length());
            const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);
            buffer.insert(buffer.end(), len_bytes, len_bytes + 4);
            const uint8_t* str_bytes = reinterpret_cast<const uint8_t*>(value.c_str());
            buffer.insert(buffer.end(), str_bytes, str_bytes + len);
        }
        return buffer;
    }

    std::string encodeJson() const {
        json j;
        j["state"] = static_cast<uint64_t>(state);
        j["value"] = value;
        return j.dump();
    }
};

class ProtocolServerCommand {
public:
    uint16_t command;
    std::string value;

    ProtocolServerCommand() {
        // All members are default initialized
    }

    void decode(const uint8_t* data, size_t data_len) {
        if (data == nullptr || data_len == 0) return;
        
        // Check if JSON format
        if (data[0] == '{') {
            try {
                std::string json_str(reinterpret_cast<const char*>(data), data_len);
                json j = json::parse(json_str);
                if (j.contains("command")) {
                    this->command = j["command"].get<uint64_t>();
                }
                if (j.contains("value")) {
                    this->value = j["value"].get<std::string>();
                }
            } catch (const std::exception& e) {
                std::cerr << "JSON parse error: " << e.what() << std::endl;
            }
            return;
        }
        
        // Binary decoding
        size_t pointer = 0;
        if (pointer + 2 <= data_len) {
            command = *reinterpret_cast<const uint16_t*>(data + pointer);
            pointer += 2;
        }
        if (pointer + 4 <= data_len) {
            int32_t value_len = *reinterpret_cast<const int32_t*>(data + pointer);
            pointer += 4;
            if (pointer + value_len <= data_len) {
                value = std::string(reinterpret_cast<const char*>(data + pointer), value_len);
                pointer += value_len;
            }
        }
    }

    std::vector<uint8_t> encodeBinary() const {
        std::vector<uint8_t> buffer;
        // Encode command
        {
            const uint16_t* ptr = &command;
            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);
            buffer.insert(buffer.end(), bytes, bytes + 2);
        }
        // Encode value
        {
            int32_t len = static_cast<int32_t>(value.length());
            const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);
            buffer.insert(buffer.end(), len_bytes, len_bytes + 4);
            const uint8_t* str_bytes = reinterpret_cast<const uint8_t*>(value.c_str());
            buffer.insert(buffer.end(), str_bytes, str_bytes + len);
        }
        return buffer;
    }

    std::string encodeJson() const {
        json j;
        j["command"] = static_cast<uint64_t>(command);
        j["value"] = value;
        return j.dump();
    }
};

#endif // PROTOCOL_H
