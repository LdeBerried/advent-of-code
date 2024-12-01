def get_locations() -> tuple[list[int], list[int]]:
    """
    Retrieve int lists from plain text file
    :return: 2 int lists ordered from lowest to highest
    """
    first_list = []
    second_list = []
    with open("location_list", "r") as location_file:
        for line in location_file:
            locations = line.split()
            first_list.append(int(locations[0]))
            second_list.append(int(locations[1]))

    first_list.sort()
    second_list.sort()
    return first_list, second_list


def get_location_distances_from_lists(first_location_list: list[int], second_location_list: list[int]) -> int:
    """
    Given 2 int lists with locations, return the sum of distance between same ordered items in both lists so the result is always a positive number since it is a distance.
    i.e.: (first_element_list_a = 3, first_element_list_b = 2, distance = 1), (second_element_list_a = 3, second_element_list_b = 6, distance = 3)
    :param first_location_list: integer list with n elements
    :param second_location_list: integer list with n elements
    :return: sum of distances between elements in one list against the other
    """
    location_distance = 0
    for first_location, second_location in zip(first_location_list, second_location_list):
        if first_location > second_location:
            location_distance += first_location - second_location
        elif second_location > first_location:
            location_distance += second_location - first_location

    return location_distance


def main():
    first_location_list, second_location_list = get_locations()
    location_distances = get_location_distances_from_lists(first_location_list, second_location_list)
    print(location_distances)


if __name__ == "__main__":
    main()
