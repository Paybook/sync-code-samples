

# Widget

### Introducción

El Widget es un recurso de Paybook Sync que asiste al usuario final en la creación de las credenciales.



### Ventajas

* El Widget implementa las mejores prácticas de seguridad para el manejo de las credenciales del usuario.
* Al añadir el Widget a tu proyecto reduces el tiempo de la integración.
* Si un sitio presenta un nuevo elemento Multi-factor para la autenticación, éste será añadido automáticamente en el Widget.  



### Elementos que debes definir

<table>
    <thead>
        <tr>
            <th>Nombre</th>
            <th>Variable en Ejemplo</th>
            <th>Descripción</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Token de Sesión</td>
            <td>sync_token_session</td>
            <td>Es el token de sesión generado en el Back End</td>
        </tr>
        <tr>
            <td>Opciones</td>
            <td>widget_options</td>
            <td>Opciones de configuración para el Widget</td>
        </tr>
        <tr>
            <td>Callback</td>
            <td>widget_callback</td>
            <td>Función que ejecuta el Widget una vez que las credenciales han sido creadas, ésto le permite al usuario retomar el flujo una vez concluída la participación del Widget</td>
        </tr>
    </tbody>
</table>



### Integración del Widget

Para agregar el widget de Paybook en tu proyecto, primero tienes que incluir un div con **id** _sync_container_:

```html
<div id="sync_container"></div>
```

Después de eso, agrega el siguiente script:

```html
<script> 
  !function(w,d,s,id,r){
    var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?"http":"https";
    if(!d.getElementById(id)){
      w[r]={};
      w[r]=w[r]||function(){w[r].q=w[r].q||[].push(arguments)};
      js=d.createElement(s);
      js.id=id;
      js.type = 'text/javascript';
      js.src=p+"://www.paybook.com/sync/widget.js";
      fjs.parentNode.insertBefore(js,fjs);
    }
  }(window,document,"script","sync-widget", "syncWidget");

  var sync_token_session = "{TOKEN}";
  var widget_callback = function(response){
    // Developer's code
  }
  var widget_options = {
    token: sync_token_session, 
    baseDiv: "sync_container", 
    theme: "light", 
    avoidAdmin: true,
    callback : widget_callback    
  }
  
  syncWidget.options = widget_options;
	if (typeof syncWidget.setToken === "function") {syncWidget.setToken(sync_token_session)} 	
</script> 
```



### Opciones del Widget

<table>
    <thead>
        <tr>
            <th>Opción</th>
            <th>Tipo</th>
            <th>Descripción</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>noLogo</td>
            <td>Booleano</td>
            <td>Cuando es verdadero, oculta los logs de los bancos y sólo muestra el texto con el nombre</td>
        </tr>
        <tr>
            <td>quickAnswer</td>
            <td>Booleano</td>
            <td>Manda un mensaje de éxito cuando el proceso ha logrado iniciar sesión</td>
        </tr>
        <tr>
            <td>exc</td>
            <td>Arreglo</td>
            <td>Arreglo de id_site_organization de las instituciones que serán excluidas.</td>
        </tr>
      	<tr>
            <td>inc</td>
            <td>Arreglo</td>
            <td>Arreglo de id_site_organization de las únicas instituciones que serán incluidas.</td>
        </tr>
      	<tr>
            <td>callback</td>
            <td>Función</td>
            <td>Callback que se llama cuando algunas acciones son ejecutadas con éxito. Regresa un objeto con las siguientes variables:
credentials: Objeto que contiene: id_credential: ID de la credencial, status: URL para obtener cambios en el estatus del proceso.
message: Mensaje de respuesta.
organizationName: Nombre de la organización.
response: Código de respuesta.
siteName: Nombre del sitio.
</td>
        </tr>
      	<tr>
            <td>theme</td>
            <td>String</td>
            <td>Cambia el tema. Actualmente los valores pueden ser "dark" o "light"</td>
        </tr>
      	<tr>
            <td>userTest</td>
            <td>Booleano</td>
            <td>Cuando el usuario necesita asistencia del equipo de Sync, es necesario tener ésta bandera activa para la depuración.</td>
        </tr>
      	<tr>
            <td>baseDiv</td>
            <td>String</td>
            <td>Id del div contenedor</td>
        </tr>
      	<tr>
            <td>widgetId</td>
            <td>String</td>
            <td>Id del widget. Es útil cuando más de un widget es requerido en la misma página.</td>
        </tr>
      	<tr>
            <td>token</td>
            <td>String</td>
            <td>Token de la sesión.</td>
        </tr>
     		<tr>
            <td>sandbox</td>
            <td>String</td>
            <td>Fija el modo de depuración.</td>
        </tr>
       	<tr>
            <td>locale</td>
            <td>String</td>
            <td>Locale define el lenguaje del usuario.</td>
        </tr>
     		<tr>
            <td>start</td>
            <td>String</td>
            <td>Si pones 'admin' el widget inicia en la lista de cuentas agregadas o inicia en un banco específico si pones el nombre del banco.</td>
        </tr>
       	<tr>
            <td>avoidAdmin</td>
            <td>Booleano</td>
            <td>Si es "true", no te permitirá ir a la lista de cuentas agregadas después de agregar una cuenta.</td>
        </tr>
     		<tr>
            <td>catalogue</td>
            <td>Arreglo</td>
            <td>Filtra por tipo de servicio, las opciones son: 'utility’, 'government’ y 'banks', para utilidades, gobierno y bancos respectivamente</td>
        </tr>
      	<tr>
            <td>siteType</td>
            <td>Arreglo</td>
            <td>Filtra por tipo de siteType, las opciones son: 'credentials’, 'document’</td>
        </tr>
    </tbody>
</table>



### Ejemplo

```html
<html>
<head>
  <script> 
    !function(w,d,s,id,r){
      var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?"http":"https";
      if(!d.getElementById(id)){
        w[r]={};
        w[r]=w[r]||function(){w[r].q=w[r].q||[].push(arguments)};
        js=d.createElement(s);
        js.id=id;
        js.type = 'text/javascript';
        js.src=p+"://www.paybook.com/sync/widget.js";
        fjs.parentNode.insertBefore(js,fjs);
      }
    }(window,document,"script","sync-widget", "syncWidget");

    var sync_token_session = "{TOKEN}";
    var widget_callback = function(response){
      console.log(response);
    }
    var widget_options = {
      token: sync_token_session, 
      baseDiv: "sync_container", 
      theme: "light", 
      avoidAdmin: true,
      callback : widget_callback    
    }

    syncWidget.options = widget_options;
    if (typeof syncWidget.setToken === "function") {syncWidget.setToken(sync_token_session)} 	
  </script> 
</head>
	<div id="sync_container" style="height: 800px;width: 600px;"></div>
</html>
```



##Casos de uso

###Configuración del Widget para iniciar en un sitio en específico:

Suponiendo que se desea iniciar en el sitio del SAT para descarga de facturas (id_site : )

```javascript
var widget_options = {
      token: sync_token_session, 
      baseDiv: "sync_container", 
      theme: "light", 
      avoidAdmin: true,
      callback : widget_callback,
  		inc : ["56cf5728784806f72b8b456f"]
    }
```

