gcc -c -o ./src/main.o ./src/main.c -I".\deps\freeglut\include"
gcc -o ./bin/main.exe ./src/main.o -L".\deps\freeglut\lib" -lglu32 -lopengl32 -lfreeglut
./bin/main.exe