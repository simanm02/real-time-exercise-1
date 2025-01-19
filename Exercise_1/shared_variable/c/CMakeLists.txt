cmake_minimum_required(VERSION 3.5)

project(AAAA C)

find_package(Threads REQUIRED)

add_executable(foo foo.c) # <<< ezt kell átírni a fájlokra
target_link_libraries(foo Threads::Threads)
# valahogy meg lehetne oldani, de egyszerűbb így
# understandable
# akkor hello

# literally 9 sor ... de az kellett hozzá, de a make-el ellentétben ez legalább tényleg cross-platform
# nem csak gányolás és imádkozás, mint a make 90 %-a, igen lmao
# igen a make nem igazán megy windowson, lehet próbálkozni, de általában nem éri meg max ha külön makefile van minden platformra
# lehet valakinnek valahol ment, de ez lehet 10 éve volt már
# igazán izgalmas lehet

# igazából h másik ilyen feladat van ezt a fájlt tudod copy-pastelni csak a
