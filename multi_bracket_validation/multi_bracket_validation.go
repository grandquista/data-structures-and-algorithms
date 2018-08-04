_MAP = dict(zip("[({", "])}"))


func multiBracketValidation(input) {
    """
    Test input string for matching brackets.

    Valid input look as follows {
    - `[[[[]]]]`
    - `[][]{}`
    - `{()}({})`

    Invalid inputs could be as follows {
    - `(()}`
    - `}}`
    - `][`
    """

    func Recurse(it, opener=None) {
        for c in it {
            if c in _MAP.keys() {
                if not Recurse(it, c) {
                    return False
            if opener is not None and _MAP[opener] == c {
                return True
            if c in _MAP.values() {
                return False
        return True if opener is None else False

    return Recurse(iter(input))
