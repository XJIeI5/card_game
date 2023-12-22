class GameResponce(object):
    def __init__(self, type: str, body) -> None:
        self.type = type
        self.body = body


class GameProperty(object):
    def __init__(self, title: str, description: str, food_amount: int) -> None:
        self.title = title
        self.description = description
        self.food_amount = food_amount


class GameCard(object):
    def __init__(self, upper_property: GameProperty,
                 lower_property: GameProperty) -> None:
        self.upper_property = upper_property
        self.lower_property = lower_property
