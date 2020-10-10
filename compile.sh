#Compiling
g++ -c -o ./objects/game.o ./common/game.cpp -I".\include" -I".\common"
g++ -c -o ./objects/resource_manager.o ./common/resource_manager.cpp -I".\include" -I".\common"
g++ -c -o ./objects/shader.o ./common/shader.cpp -I".\include" -I".\common"
g++ -c -o ./objects/texture.o ./common/texture.cpp -I".\include" -I".\common"
g++ -c -o ./objects/main.o ./src/main.cpp -I".\include" -I".\common"
g++ -c -o ./objects/sprite_renderer.o ./common/sprite_renderer.cpp -I".\include" -I".\common"

#Linking
g++ -o ./bin/main.exe ./objects/game.o ./objects/sprite_renderer.o ./objects/resource_manager.o ./objects/shader.o ./objects/texture.o ./objects/main.o -L".\libraries\glew" -L".\libraries\glfw" -lopengl32 -lglew32 -lglfw3

#Running
echo Running Program Now...
./bin/main.exe
echo Closing Now...
echo ...Goodbye!