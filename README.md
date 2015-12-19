# srcgen
generate source code from data

## Template

Uses standard go templates with the default delimiters set to <% %>, this can be changed on a per template basis.

## Creating a Template

Templates is stored in the template folder, where a folder specify the template name.

Example:
We have a template named 'readme' the folder structure should look like this.

    ./TemplatePath/readme/readme.md

By default srcgen will traverse all files in the folder and execute them as templates, how ever more complex behavior can be defined using a template settings file.


## Template Settings

## Template Functions

srcgen uses the Go html/template library for its template engine.
It is an extremely lightweight engine that provides a very small amount of
logic.

Go templates are lightweight but extensible. The following
functions was added to the basic template logic.

(Go itself supplies built-in functions, including comparison operators
and other basic tools; these are listed in the
[Go template documentation](http://golang.org/pkg/text/template/#hdr-Functions).)

## General Template Functions

### delimit
Loops through any array, slice or map and returns a string of all the values separated by the delimiter. There is an optional third parameter that lets you choose a different delimiter to go between the last two values.
Maps will be sorted by the keys, and only a slice of the values will be returned, keeping a consistent output order.

Works on [lists](/templates/list/), [taxonomies](/taxonomies/displaying/), [terms](/templates/terms/), [groups](/templates/list/)

e.g.

    // Front matter
    +++
    tags: [ "tag1", "tag2", "tag3" ]
    +++

    // Used anywhere in a template
    Tags: {{ delimit .Params.tags ", " }}

    // Outputs Tags: tag1, tag2, tag3

    // Example with the optional "last" parameter
    Tags: {{ delimit .Params.tags ", " " and " }}

    // Outputs Tags: tag1, tag2 and tag3

### dict
Creates a dictionary (map[string, interface{}), expects parameters added in value:object fasion.
Invalid combinations like keys that are not strings or uneven number of parameters, will result in an exception thrown
Useful for passing maps to partials when adding to a template.

e.g. Pass into "foo.html" a map with the keys "important, content"

    {{$important := .Site.Params.SomethingImportant }}
    {{range .Site.Params.Bar}}
        {{partial "foo" (dict "content" . "important" $important)}}
    {{end}}

"foo.html"

    Important {{.important}}
    {{.content}}


or Create a map on the fly to pass into

    {{partial "foo" (dict "important" "Smiles" "content" "You should do more")}}



### echoParam
Prints a parameter if it is set.

e.g. `{{ echoParam .Params "project_url" }}`


### eq
Returns true if the parameters are equal.

e.g.

    {{ if eq .Section "blog" }}current{{ end }}


### first
Slices an array to only the first _N_ elements.

Works on [lists](/templates/list/), [taxonomies](/taxonomies/displaying/), [terms](/templates/terms/), [groups](/templates/list/)

e.g.

    {{ range first 10 .Data.Pages }}
        {{ .Render "summary" }}
    {{ end }}

### last
Slices an array to only the last _N_ elements.

Works on [lists](/templates/list/), [taxonomies](/taxonomies/displaying/), [terms](/templates/terms/), [groups](/templates/list/)

e.g.

    {{ range last 10 .Data.Pages }}
        {{ .Render "summary" }}
    {{ end }}

### after
Slices an array to only the items after the <em>N</em>th item. Use this in combination
with `first` to use both halves of an array split at item _N_.

Works on [lists](/templates/list/), [taxonomies](/taxonomies/displaying/), [terms](/templates/terms/), [groups](/templates/list/)

e.g.

    {{ range after 10 .Data.Pages }}
        {{ .Render "title" }}
    {{ end }}

### getenv
Returns the value of an environment variable.

Takes a string containing the name of the variable as input. Returns
an empty string if the variable is not set, otherwise returns the
value of the variable. Note that in Unix-like environments, the
variable must also be exported in order to be seen by `hugo`.

e.g.

    {{ getenv "HOME" }}


### in
Checks if an element is in an array (or slice) and returns a boolean.
The elements supported are strings, integers and floats (only float64 will match as expected).
In addition, it can also check if a substring exists in a string.

e.g.

    {{ if in .Params.tags "Git" }}Follow me on GitHub!{{ end }}

or

    {{ if in "this string contains a substring" "substring" }}Substring found!{{ end }}


### intersect
Given two arrays (or slices), this function will return the common elements in the arrays.
The elements supported are strings, integers and floats (only float64).

A useful example of this functionality is a 'similar posts' block.
Create a list of links to posts where any of the tags in the current post match any tags in other posts.

e.g.

    <ul>
    {{ $page_link := .Permalink }}
    {{ $tags := .Params.tags }}
    {{ range .Site.Pages }}
        {{ $page := . }}
        {{ $has_common_tags := intersect $tags .Params.tags | len | lt 0 }}
        {{ if and $has_common_tags (ne $page_link $page.Permalink) }}
            <li><a href="{{ $page.Permalink }}">{{ $page.Title }}</a></li>
        {{ end }}
    {{ end }}
    </ul>


### isset
Returns true if the parameter is set.
Takes either a slice, array or channel and an index or a map and a key as input.

e.g. `{{ if isset .Params "project_url" }} {{ index .Params "project_url" }}{{ end }}`

### seq - not supported

Creates a sequence of integers. It's named and used as GNU's seq.

Some examples:

* `3` => `1, 2, 3`
* `1 2 4` => `1, 3`
* `-3` => `-1, -2, -3`
* `1 4` => `1, 2, 3, 4`
* `1 -2` => `1, 0, -1, -2`

### sort
Sorts maps, arrays and slices, returning a sorted slice.
A sorted array of map values will be returned, with the keys eliminated.
There are two optional arguments, which are `sortByField` and `sortAsc`.
If left blank, sort will sort by keys (for maps) in ascending order.

Works on [lists](/templates/list/), [taxonomies](/taxonomies/displaying/), [terms](/templates/terms/), [groups](/templates/list/)

e.g.

    // Front matter
    +++
    tags: [ "tag3", "tag1", "tag2" ]
    +++

    // Site config
    +++
    [params.authors]
      [params.authors.Derek]
        "firstName"  = "Derek"
        "lastName"   = "Perkins"
      [params.authors.Joe]
        "firstName"  = "Joe"
        "lastName"   = "Bergevin"
      [params.authors.Tanner]
        "firstName"  = "Tanner"
        "lastName"   = "Linsley"
    +++

    // Use default sort options - sort by key / ascending
    Tags: {{ range sort .Params.tags }}{{ . }} {{ end }}

    // Outputs Tags: tag1 tag2 tag3

    // Sort by value / descending
    Tags: {{ range sort .Params.tags "value" "desc" }}{{ . }} {{ end }}

    // Outputs Tags: tag3 tag2 tag1

    // Use default sort options - sort by value / descending
    Authors: {{ range sort .Site.Params.authors }}{{ .firstName }} {{ end }}

    // Outputs Authors: Derek Joe Tanner

    // Use default sort options - sort by value / descending
    Authors: {{ range sort .Site.Params.authors "lastName" "desc" }}{{ .lastName }} {{ end }}

    // Outputs Authors: Perkins Linsley Bergevin


### where
Filters an array to only elements containing a matching value for a given field.

Works on [lists](/templates/list/), [taxonomies](/taxonomies/displaying/), [terms](/templates/terms/), [groups](/templates/list/)

e.g.

    {{ range where .Data.Pages "Section" "post" }}
       {{ .Content }}
    {{ end }}

It can be used with dot chaining second argument to refer a nested element of a value.

e.g.

    // Front matter on some pages
    +++
    series: golang
    +++

    {{ range where .Site.Pages "Params.series" "golang" }}
       {{ .Content }}
    {{ end }}

It can also be used with an operator like `!=`, `>=`, `in` etc. Without an operator (like above), `where` compares a given field with a matching value in a way like `=` is specified.

e.g.

    {{ range where .Data.Pages "Section" "!=" "post" }}
       {{ .Content }}
    {{ end }}

Following operators are now available

- `=`, `==`, `eq`: True if a given field value equals a matching value
- `!=`, `<>`, `ne`: True if a given field value doesn't equal a matching value
- `>=`, `ge`: True if a given field value is greater than or equal to a matching value
- `>`, `gt`: True if a given field value is greater than a matching value
- `<=`, `le`: True if a given field value is lesser than or equal to a matching value
- `<`, `lt`: True if a given field value is lesser than a matching value
- `in`: True if a given field value is included in a matching value. A matching value must be an array or a slice
- `not in`: True if a given field value isn't included in a matching value. A matching value must be an array or a slice

*`where` and `first` can be stacked, e.g.:*

    {{ range first 5 (where .Data.Pages "Section" "post") }}
       {{ .Content }}
    {{ end }}

### Unset field
Filter only work for set fields. To check whether a field is set or exist, use operand `nil`.

This can be useful to filter a small amount of pages from a large pool. Instead of set field on all pages, you can set field on required pages only.

Only following operators are available for `nil`

- `=`, `==`, `eq`: True if the given field is not set.
- `!=`, `<>`, `ne`: True if the given field is set.

e.g.

    {{ range where .Data.Pages ".Params.specialpost" "!=" nil }}
       {{ .Content }}
    {{ end }}


## Math

<table class="table table-bordered">
<thead>
<tr>
<th>Function</th>
<th>Description</th>
<th>Example</th>
</tr>
</thead>

<tbody>
<tr>
<td><code>add</code></td>
<td>Adds two integers.</td>
<td><code>{{add 1 2}}</code> → 3</td>
</tr>

<tr>
<td><code>div</code></td>
<td>Divides two integers.</td>
<td><code>{{div 6 3}}</code> → 2</td>
</tr>

<tr>
<td><code>mod</code></td>
<td>Modulus of two integers.</td>
<td><code>{{mod 15 3}}</code> → 0</td>
</tr>

<tr>
<td><code>modBool</code></td>
<td>Boolean of modulus of two integers.  <code>true</code> if modulus is 0.</td>
<td><code>{{modBool 15 3}}</code> → true</td>
</tr>

<tr>
<td><code>mul</code></td>
<td>Multiplies two integers.</td>
<td><code>{{mul 2 3}}</code> → 6</td>
</tr>

<tr>
<td><code>sub</code></td>
<td>Subtracts two integers.</td>
<td><code>{{sub 3 2}}</code> → 1</td>
</tr>

</tbody>
</table>


## Strings

### chomp
Removes any trailing newline characters. Useful in a pipeline to remove newlines added by other processing (including `markdownify`).

e.g., `{{chomp "<p>Blockhead</p>\n"}}` → `"<p>Blockhead</p>"`


### lower
Converts all characters in string to lowercase.

e.g. `{{lower "BatMan"}}` → "batman"

### pluralize
Pluralize the given word with a set of common English pluralization rules.

e.g. `{{ "cat" | pluralize }}` → "cats"

### replace
Replaces all occurrences of the search string with the replacement string.

e.g. `{{ replace "Batman and Robin" "Robin" "Catwoman" }}` → "Batman and Catwoman"

### singularize
Singularize the given word with a set of common English singularization rules.

e.g. `{{ "cats" | singularize }}` → "cat"

### slicestr

Slicing in `slicestr` is done by specifying a half-open range with two indices, `start` and `end`.
For example, 1 and 4 creates a slice including elements 1 through 3.
The `end` index can be omitted; it defaults to the string's length.

e.g.

* `{{slicestr "BatMan" 3}}` → "Man"
* `{{slicestr "BatMan" 0 3}}` → "Bat"

### substr

Extracts parts of a string, beginning at the character at the specified
position, and returns the specified number of characters.

It normally takes two parameters: `start` and `length`.
It can also take one parameter: `start`, i.e. `length` is omitted, in which case
the substring starting from start until the end of the string will be returned.

To extract characters from the end of the string, use a negative start number.

In addition, borrowing from the extended behavior described at http://php.net/substr,
if `length` is given and is negative, then that many characters will be omitted from
the end of string.

e.g.

* `{{substr "BatMan" 0 -3}}` → "Bat"
* `{{substr "BatMan" 3 3}}` → "Man"

### title
Converts all characters in string to titlecase.

e.g. `{{title "BatMan"}}` → "Batman"


### trim
Returns a slice of the string with all leading and trailing characters contained in cutset removed.

e.g. `{{ trim "++Batman--" "+-" }}` → "Batman"


### upper
Converts all characters in string to uppercase.

e.g. `{{upper "BatMan"}}` → "BATMAN"




## Content Views

### Render
Takes a view to render the content with.  The view is an alternate layout, and should be a file name that points to a template in one of the locations specified in the documentation for [Content Views](/templates/views).

This function is only available on a piece of content, and in list context.

This example could render a piece of content using the content view located at `/layouts/_default/summary.html`:

    {{ range .Data.Pages }}
        {{ .Render "summary"}}
    {{ end }}



## Advanced

### apply

Given a map, array, or slice, returns a new slice with a function applied over it. Expects at least three parameters, depending on the function being applied. The first parameter is the sequence to operate on; the second is the name of the function as a string, which must be in the Hugo function map (generally, it is these functions documented here). After that, the parameters to the applied function are provided, with the string `"."` standing in for each element of the sequence the function is to be applied against. An example is in order:

    +++
    names: [ "Derek Perkins", "Joe Bergevin", "Tanner Linsley" ]
    +++

    {{ apply .Params.names "urlize" "." }} → [ "derek-perkins", "joe-bergevin", "tanner-linsley" ]

This is roughly equivalent to:

    {{ range .Params.names }}{{ . | urlize }}{{ end }}

However, it isn’t possible to provide the output of a range to the `delimit` function, so you need to `apply` it. A more complete example should explain this. Let's say you have two partials for displaying tag links in a post,  "post/tag/list.html" and "post/tag/link.html", as shown below.

    <!-- post/tag/list.html -->
    {{ with .Params.tags }}
    <div class="tags-list">
      Tags:
      {{ $len := len . }}
      {{ if eq $len 1 }}
        {{ partial "post/tag/link" (index . 0) }}
      {{ else }}
        {{ $last := sub $len 1 }}
        {{ range first $last . }}
          {{ partial "post/tag/link" . }},
        {{ end }}
        {{ partial "post/tag/link" (index . $last) }}
      {{ end }}
    </div>
    {{ end }}


    <!-- post/tag/link.html -->
    <a class="post-tag post-tag-{{ . | urlize }}" href="/tags/{{ . | urlize }}">{{ . }}</a>

This works, but the complexity of "post/tag/list.html" is fairly high; the Hugo template needs to perform special behaviour for the case where there’s only one tag, and it has to treat the last tag as special. Additionally, the tag list will be rendered something like "Tags: tag1 , tag2 , tag3" because of the way that the HTML is generated and it is interpreted by a browser.

This is Hugo. We have a better way. If this were your "post/tag/list.html" instead, all of those problems are fixed automatically (this first version separates all of the operations for ease of reading; the combined version will be shown after the explanation).

    <!-- post/tag/list.html -->
    {{ with.Params.tags }}
    <div class="tags-list">
      Tags:
      {{ $sort := sort . }}
      {{ $links := apply $sort "partial" "post/tag/link" "." }}
      {{ $clean := apply $links "chomp" "." }}
      {{ delimit $clean ", " }}
    </div>
    {{ end }}

In this version, we are now sorting the tags, converting them to links with "post/tag/link.html", cleaning off stray newlines, and joining them together in a delimited list for presentation. That can also be written as:

    <!-- post/tag/list.html -->
    {{ with.Params.tags }}
    <div class="tags-list">
      Tags:
      {{ delimit (apply (apply (sort .) "partial" "post/tag/link" ".") "chomp" ".") ", " }}
    </div>
    {{ end }}

`apply` does not work when receiving the sequence as an argument through a pipeline.

***

### base64Encode and base64Decode

`base64Encode` and `base64Decode` let you easily decode content with a base64 encoding and vice versa through pipes. Let's take a look at an example:


    {{ "Hello world" | base64Encode }}
    <!-- will output "SGVsbG8gd29ybGQ=" and -->

    {{ "SGVsbG8gd29ybGQ=" | base64Decode }}
    <!-- becomes "Hello world" again. -->

You can also pass other datatypes as argument to the template function which tries
to convert them. Now we use an integer instead of a string:


    {{ 42 | base64Encode | base64Decode }}
    <!-- will output "42". Both functions always return a string. -->

**Tip:** Using base64 to decode and encode becomes really powerful if we have to handle
responses of APIs.

    {{ $resp := getJSON "https://api.github.com/repos/spf13/hugo/readme"  }}
    {{ $resp.content | base64Decode | markdownify }}

The response of the GitHub API contains the base64-encoded version of the [README.md](https://github.com/spf13/hugo/blob/master/README.md) in the Hugo repository. Now we can decode it and parse the Markdown. The final output will look similar to the rendered version on GitHub.
