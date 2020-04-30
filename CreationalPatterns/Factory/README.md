#Factory Pattern or Factory method Pattern
    Factory method is most commonly used creational design pattern.
    This is useful when there are different kind of objects whose creation logic varies for each type and these objects serve same business logic for the client (a family of related objects) also choice of selection of these objects is parameters dependent.
##Scenario:
    When Logic for construction of object is chosen depending on certain perameters then such logic can be abstracted called as factory method, that takes parameters to choose which logic to be performed to generate an object.
    Tobe noted that sometimes a logic is chosen and sometimes a type is directly selected from the parameters provided point here to be noted is that products are generated depending on parameters.
    
##UML Diagram
![Factory method Design pattern](Factory_Method.png)
[pic source](https://sourcemaking.com/design_patterns/factory_method)