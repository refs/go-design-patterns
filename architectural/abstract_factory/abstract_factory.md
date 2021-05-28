
/*
- The abstract factory pattern provides a way to encapsulate a group of individual factories that have a common theme without specifying their concrete classes.
- The client software creates a concrete implementation of the abstract factory and then uses the generic interface of the factory to create the concrete objects.
- The client does not know (or care) which concrete objects it gets from each of these internal factories, since it uses only the generic interfaces of their products.
*/

// Abstract factory. It determines the concrete type to be created. Insulates client code from object creation by having
// clients ask a factory object to create an object of the desired abstract type and to return an abstract pointer to
// the object.