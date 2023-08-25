##Parte 2

En esta carpeta del proyecto se encuentra la segunda parte de la prueba técnica. Esta parte está hecha en Go. Para ejecutarla es necesario clonar el repositorio y, estando ubicados en el directorio *parte2*, correr los el siguiente comando:

```go run main.go ```

En la prueba no se especificaba la forma en la que se accedía al JSON generado por el programa, así que en este caso asumí que también se podía hacer un servicio web (al igual que en la parte 1), para recibir el JSON a través de una petición GET. Por defecto, el proyecto corre en el puerto 8080, por lo que la URL raíz es http://localhost:8080. Esta API solamente tiene un endpoint, que es el que se encarga de resolver la tarea encomendada en la prueba. Se trata de una petición GET con la URL http://localhost:8080/organize-csv. Al acceder a esta URL (si el proyecto está corriendo), se podrán ver los resultados del JSON solicitado.

El archivo csv de entrada debe estar en la misma carpeta del archivo de código fuente y debe llamarse *file.csv*. En el repositorio hay un archivo de ejemplo. Este archivo puede ser modificado, pero para que los cambios se vean reflejados en el endpoint, es necesario detener la ejecución y volver a correr el comando especificado al principio. 