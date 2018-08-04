from .multi_bracket_validation import multi_bracket_validation


def test_valid_input():
    assert multi_bracket_validation("[hello world]")


def test_inverted_brackets():
    assert not multi_bracket_validation("]hello world[")


def test_only_close():
    assert not multi_bracket_validation("))))")


def test_only_open():
    assert not multi_bracket_validation("{{{{")


def test_deep_nesting():
    assert multi_bracket_validation("{" * 101 + "}" * 101)


def test_repeated_groups():
    assert multi_bracket_validation("({[({()})]}){{(([[]]))}}")
