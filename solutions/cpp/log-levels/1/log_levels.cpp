#include <string>

namespace log_line {
std::string message(std::string line) {
    size_t idx = line.find(':');
    return line.substr(idx+2);
}

std::string log_level(std::string line) {
    size_t start = line.find('[') + 1;
    size_t end = line.find(']');
    return line.substr(start,end - start);
}

std::string reformat(std::string line) {
  return message(line) + " (" + log_level(line) + ")";
}
}  // namespace log_line
