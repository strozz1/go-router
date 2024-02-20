# Dynamic Routes

pd: esto son notas para el futuro oscar, espero que seas mejor porque madre mia que spaggueti code

## Primera version
Acabo de terminar de implementar la primera version de lo que probalbemente cambie en un futuro. He implementado las rutas dinamicas y voy a explicar
mi procedimiento.

La idea que tenia era cuando un usuario registra un endpoint, y quiere que sea dinamico, debe rodearlo con corchetes `/{id}`. Debe de acabar
en una letra. Y haciendo uso del atributo segmentType ponemos el segmento a dinamico.

Lo primero que se verifica es que la ruta dinamica tenga un formato correcto (que no tenga varios guiones,sin espacios...) y esto se verifica
con un regex.
Hay 3 metodos relacionados con regex.

1. Uno verifica si se trata de una ruta dinamica, simplemente mira si es del tipo `/{x}`.
2. Otro verifica si la sintaxis de la ruta dinamica es correcta.
3. El ultimo es para  rutas estaticas.

De esta forma, ya tenemos una manera de que el usuario cree rutas dinamicas.

A la hora de leer una peticion, el flujo sigue siendo el mismo solo que cuando leemos un segmento, primero miramos si el segmento es dinamico.
Si no lo es simplemente lo tratamos como anteriormente. Si se trata de uno dinamico,si lo piensas relmente, cualquier valor que encontremos en el
segmento es valido (ya que es dinamico). Otra cosa seria si por ejemplo el id usuario es valido o existe. Pero eso son problemas relacionados
con el uso que se le da a la ruta.

Por lo que, si es dinamica, practicamente es una pase valido para pasar al siguiente nivel.

### Fallos

Asi analizando un poco el codigo, veo que cuando el usuario quiere obtener el valor dinamico, es un poco engorroso porque lo va a tener que escribir el.
Asique para un futuro hay que pensar en como se podria solucionar esto. De momento da igual XD.

### Conclusiones

Mi codigo no es el mejor ni de cerca y tiene muchos puntos de mejora. Pero para una primera version que funciona y detecta si la ruta esta bien formada me parece suficiente.
El mejorar el codigo y su lectura sera para otro momento.
