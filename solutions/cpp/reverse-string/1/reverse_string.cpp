#include "reverse_string.h"

namespace reverse_string {

    std::string reverse_string(const std::string& input) {
    std::string result;

    for (int i = input.length() - 1; i >= 0; --i) {
        result += input[i];
    }

    return result;
}

}  // namespace reverse_string
