from typing import Callable
import pkg.values as v


class Property:
    def __init__(self) -> None:
        self.is_unique: bool = False


class StatusProperty(Property):
    def __init__(self, action: v.StatusAction, duration: int) -> None:
        super().__init__()
        self._action: v.StatusAction = action
        self.duration: int = duration
    
    def get_status_action(self) -> v.StatusAction:
        return self._action


class DamageProperty(Property):
    def __init__(self, damage: v.Damage) -> None:
        super().__init__()
        self._damage: v.Damage = damage
    
    def get_damage(self) -> v.Damage:
        return self._damage
