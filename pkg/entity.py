from typing import List
import pkg.property as prop
import pkg.values as v


class Entity:
    def __init__(self) -> None:
        self.status_properties: List[prop.StatusProperty] = []
        self.damage_properties: List[prop.DamageProperty] = []
        self.unique_properties: List[prop.Property] = []

        self._hp: int = 1
    
    def add_properties(self, properties: List[prop.Property]) -> bool:
        for property in properties:
            if not self.can_add(property):
                return False
        
        for property in properties:
            if property.is_unique:
                self.unique_properties.append(property)
                continue
        
            list_corresponding_to_prop_type = {
                prop.StatusProperty: self.status_properties,
                prop.DamageProperty: self.damage_properties
            }
            for inst_class, corresponding_list in list_corresponding_to_prop_type.items():
                if isinstance(property, inst_class):
                    corresponding_list.append(property)
                    break
        return True
    
    def can_add(self, property: prop.Property) -> bool:
        return not (property in self.unique_properties)
    
    def take_damage(self, damage_props: List[prop.DamageProperty]) -> None:
        for prop in damage_props:
            self._apply_damage(prop.get_damage())

    def _apply_damage(self, damage: v.Damage) -> None:
        self._hp -= damage.value
        # TODO: death
    
    def apply_status(self) -> None:
        for status in self.status_properties:
            _apply_status_action(self, status.get_status_action())
            pass


def _apply_status_action(entity: Entity, action: v.StatusAction) -> None:
    data = action.corresponding_data
    match action.type:
        case v.StatusActionType.take_damage:
            entity._apply_damage(data)
