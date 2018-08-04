from .multiBracketValidation import multiBracketValidation


func TestValidInput() {
    assert multiBracketValidation("[hello world]")


func TestInvertedBrackets() {
    assert not multiBracketValidation("]hello world[")


func TestOnlyClose() {
    assert not multiBracketValidation("))))")


func TestOnlyOpen() {
    assert not multiBracketValidation("{{{{")


func TestDeepNesting() {
    assert multiBracketValidation("{" * 101 + "}" * 101)


func TestRepeatedGroups() {
    assert multiBracketValidation("({[({()})]}){{(([[]]))}}")
