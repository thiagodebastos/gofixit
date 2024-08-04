# Issue Tracker Written in Go

## Domain Driven Design Basics

1. **Entities and Value Objects**: Define your core domain models.
2. **Repositories**: Create interfaces for persisting and retrieving domain models.
3. **Domain Services**: Encapsulate domain logic that doesn’t naturally fit within an entity or value object.
4. **Application Services**: Coordinate tasks and orchestrate the use of domain objects and services.
5. **Factories**: Encapsulate the creation logic for complex aggregates or entities.
6. **Specifications**: Encapsulate domain logic to check if objects satisfy certain criteria.
7. **Domain Events**: Capture significant changes in the domain and react to them.
8. **Infrastructure Layer**: Handle persistence, messaging, and external system integration.
