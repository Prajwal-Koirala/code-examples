# Change an element's value at a certain index.
def change_element_at_slice_via_index(provided_slice, provided_index, provided_element):
    provided_slice[provided_index] = provided_element
    return provided_slice


def main():
    # Make a slice using a random set of data.
    randomSlice = ["a", "b", "c"]
    # Change an element's value at a certain index.
    print(change_element_at_slice_via_index(randomSlice, 1, "e"))


main()
