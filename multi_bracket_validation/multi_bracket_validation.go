_MAP = dict(zip("[({", "])}"))


def multi_bracket_validation(input):
    """
    Test input string for matching brackets.

    Valid input look as follows:
    - `[[[[]]]]`
    - `[][]{}`
    - `{()}({})`

    Invalid inputs could be as follows:
    - `(()}`
    - `}}`
    - `][`
    """

    def _recurse(it, opener=None):
        for c in it:
            if c in _MAP.keys():
                if not _recurse(it, c):
                    return False
            if opener is not None and _MAP[opener] == c:
                return True
            if c in _MAP.values():
                return False
        return True if opener is None else False

    return _recurse(iter(input))
