extern "C" {
    int writet1(char*, char*);

    int main() {
        char in[] = "hello\0 world";
        char* out;
        int ret = writet1(in, out);
        return ret;
    }
}