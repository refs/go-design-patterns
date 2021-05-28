## Software Design Patterns - With a touch of Go

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

</details>

| **Pattern** | **Description** |
|:---:| --- |
| [Factory Method](/creational/factory_method/factory.md) | Define an interface for creating an object, but let subclasses decide which class to instantiate |
| [Abstract Factory](/creational/abstract_factory/abstract_factory.md) | Encapsulate a group of individual factories |

## Sources and Inspiration
- [Software Design Patterns - Classification](https://en.wikipedia.org/wiki/Software_design_pattern#Classification_and_list)
- [tmrts/go-patterns, sadly no longer maintained](github.com/tmrts/go-patterns)
