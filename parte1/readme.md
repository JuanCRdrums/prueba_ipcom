## Parte 1

En esta carpeta del proyecto se encuentra la solución a la parte 1 de la prueba técnica. Esta parte está hecha en NodeJS. Para ejecutarla es necesario clonar el repositorio y, estando ubicados en el directorio *parte1*, correr los siguientes comandos:

``` npm install ``` para instalar todas las dependencias necesarias

``` npm start ``` para iniciar el servidor

Por defecto, el servidor corre en el puerto 3000. Por lo tanto, la ruta para acceder a la API en el navegador o en ualquier herramienta para implementar API REST (tipo Postman) es http://localhost:3000.

La API solamente tiene un endpoint, que fue el solicitado en la prueba técnica. Este endoint tiene el siguiente formato: http://localhost:3000/resumen/{YYYY-MM-DD}?dias={x}, siendo *dias* un parámetro numérico. La fecha después del slash de la URL debe ser pasada en el formato especificado: YYYY-MM-DD. Este endpoint corresponde a una petición GET. 