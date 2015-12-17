# srcgen

note! look at the swagger.io documentation!

stdgen aims to set a standard for data objects used in code generation,
the advantage of such a standard is that templates can be created by any developer
that conforms to the standard and the template will work.

the standard specify minimum functionality and can be extended by adding custom properties to
a data object. The custom properties should be well documented on the template ReadMe.md file.

## Syntax

* fields marked with [] square brackets are optional.

## Tags - Property

tags is an array of strings and is a optional property to most structures. tags
allows the developer to specify boolean type switch on objects that can be used by
filters to exclude or include a object in a templates configuration data.

## Structures

the following structure types is supported.

* interface
* model

## Application - Structure

defines a application.

* name : string
* description : string

## Models - Structure

models represent data structures that can be passed between client and server,
between function calls as arguments and stored in memory, database or file.

Typically CRUD operations is done on models but not always.

**Model structure**

* name : string
* fields : array< Field >
* [keyfield] : string
* [caption] : string
* [description] : string
* [tags] : array< string > optional


**Model.fields property**

Field type
* name : string
* type : string
* [defaultValue] : string
* [validation] : Validation Type
* [caption] : string
* [description] : string
* [tags] : array< string > optional

** Field.type property**

The 'type' field supports the following types, and can be used by your template
to convert the type into the type required by your template. any other type is
assumed to be a custom type defined in your project code.

Supported Types
* string
* decimal (f16, f32, f64)
* int (i8, i16, i32, i64)
* byte
* date
* bool
* array type
* map type type
* any

** Field.validation property**

Validation is used to check if the field value conforms to the requested data. the validation property is a simple string.

Validation Type
* required : boolean
* max : if the type is a string max refers to the strings maximum length.
* min : if the type is a string min refers to the strings minimum length.
* regex : only for string types.


## Interface

interface data contains interface call information.

**Interface Type**
* name
* functions
* [tags]

**function type**
* name : string
* arguments : [FieldType]
* returnType : string representing a std type.
