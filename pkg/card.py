from typing import List, Self

import pkg.property as prop
import pkg.entity as ent


class Card:
    def __init__(self, name: str, description: str, properties: List[prop.Property]) -> None:
        self.name: str = name
        self.description: str = description
        self.properties: List[prop.Property] = properties
    
    def __copy__(self) -> Self:
        return Card(self.name, self.description, self.properties)
    
    def play_as_creature(self) -> ent.Entity:
        return ent.Entity()

    def play_as_property(self, entity: ent.Entity) -> bool:
        return entity.add_property(self.properties)
