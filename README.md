## Software Design Patterns - with a touch of Go

### Creational

Deal with object creation mechanism; encapsulates knowledge about concrete types and how these are created.

<details>
<summary>diagram</summary>
<br>

```
┌──────────────┐ 
│              │ 
│   Creator    │ 
│              │ 
└──────────────┘ 
        △        
        │        
        │        
┌───────────────┐
│               │
│ConcreteCreator│
│               │
└───────────────┘
```

|**Pattern**|
|[Factory](https://github.com/refs/go-design-patterns/blob/master/creational/factory/factory.md)|
|[Abstract Factory](https://github.com/refs/go-design-patterns/blob/master/creational/abstract_factory/abstract_factory.md)|

</details>

## Sources and Inspiration
- [Software Design Patterns - Classification](https://en.wikipedia.org/wiki/Software_design_pattern#Classification_and_list)
- [tmrts/go-patterns, sadly no longer maintained](github.com/tmrts/go-patterns)
