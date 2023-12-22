from typing import Dict


class Request:
    def __init__(self, req_type: str, body: Dict):
        self.type = req_type
        self.body: Dict = body

    def __str__(self) -> str:
        type_str = f'"type":"{self.type}"'
        attr_list = []
        for key, value in self.body.items():
            match value:
                case bool() as bool_value:
                    curr_atr = f'"{key}":{str(bool_value).lower()}'
                case int() as int_value:
                    curr_atr = f'"{key}":{int_value}'
                case _:
                    curr_atr = f'"{key}":"{value}"'
            attr_list.append(curr_atr)
        attr_str = ','.join(attr_list)
        body_str = '"body":{' + attr_str + '}'
        return '{' + type_str + ',' + body_str + '}'
