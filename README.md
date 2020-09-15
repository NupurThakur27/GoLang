# GoLang
Practicing GoLang 
### To sort an array of structs based on a field:

```
sort.Slice(array[:], func(i, j int) bool {
		return array[i].FieldName < array[j].FieldName
})
```
