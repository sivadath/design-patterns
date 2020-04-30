#Abstract Factory
    This is an extention to Factory method pattern. To say it simple it's factory of factory methods. Another most widely used pattern.

##Scenario
    Consider a couple of objects needed in package for a client to perform some operations (certain known and limited operations). We can create an interface that is an abstraction of both of these objects also serves purpose of client software. Think how do we create objects of our compund interface we generate.
    Very first problem is that each object of this couple (generally a bundle of objects speaking only couple for simplicity) can be created differently depending on some parameters (should be reminding of factory method). Now for suppose our couple be package of object A and object B object A can be of N different types but all these types after getting created serves same purpose
    also same with B it can be of M different types, now our new interface can also be created N*M different types (generally less than or equal to N*M as in real scenarios we might not want few combinations) but once a packge of A and B is created together they serve same purpose.
    How do we create an object of new couple. Problem statement is same... i.e., depending on the combination we follow different method of creation involving creating A followed by B, so we just create a factory method that generate our couple depending on combination (perameter) passed. That's it this abstraction of multiple factories resulting in new factory method is called Abstract Factory method just like the name indicating.
    
##Example Description:
    In the example written here there are no such couple of objects but just two different factories namely cars factory and bikes factory. Yet comes under abstract factory type because it's a factory of factories.
    
    There may be a client software A that requires cars interaface (doesn't matter which car once it is created) so we created cars factory method that knows to build different types of cars. Example : Second hand car retailer website that just needs cars interface (methods like GetDoorsCount) to evaluate value.
    
    There may be another client software B that required bikes interface (doesn't matter which bike it is once created) so we created bikes factory method that knows to build different types f bikes. Example : A bike rental website that just needs bikes interface (methods like GetBikeType) to evaluate bike rent per hour.
    
    Now think that local government made tie up with both company A and B to use thier vehicles while available for police, medical e.t.c., emergencies.
    
    Now government needs a vehicle interface (doen't matter which type of vehicle once it's created) so we create a vehicles factory that knows to build both cars and bikes and once vehicle is created government sofware center uses NumberOfSeats method to evaluate how many personnals they can send at a time.
    
    Here too we created an abstraction of two different factories. 