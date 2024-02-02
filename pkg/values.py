from enum import Enum

class DamageType(Enum):
    physical = 0
    magic = 1


class Damage:
    def __init__(self, value: int, damage_type: DamageType = DamageType.physical) -> None:
        self.value: int = value
        self.type: DamageType = damage_type


class StatusActionType(Enum):
    take_damage = 0


class StatusAction:
    def __init__(self, action_type: StatusActionType, value) -> None:
        self.type: StatusActionType = action_type
        self.corresponding_data = value
