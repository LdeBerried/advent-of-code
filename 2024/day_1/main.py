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


def get_location_distance_from_lists(first_location_list: list[int], second_location_list: list[int]) -> int:
    """
    Given 2 int lists with locations, return the sum of distance between same ordered items in both lists so the result is always a positive number since it is a distance.
    i.e.: (first_element_list_a = 3, first_element_list_b = 2, distance = 1), (second_element_list_a = 3, second_element_list_b = 6, distance = 3)
    :param first_location_list: integer list with n elements
    :param second_location_list: integer list with n elements
    :return: sum of distances between elements in one list against the other
    """
    location_distance = 0
    for first_location, second_location in zip(first_location_list, second_location_list):
        location_distance += abs(first_location-second_location)
    return location_distance


def get_similarity_score(first_location_list: list[int], second_location_list: list[int]) -> int:
    similarity_score = 0
    for first_location in first_location_list:
        second_list_matches = [match for match in second_location_list if match == first_location]
        if len(second_list_matches) > 0:
            similarity_score += first_location * len(second_list_matches)
    return similarity_score


def main():
    first_location_list, second_location_list = get_locations()
    location_distance = get_location_distance_from_lists(first_location_list, second_location_list)
    print(f"Location distance between both lists is: {location_distance}")

    similarity_score = get_similarity_score(first_location_list, second_location_list)
    print(f"Similarity score between both lists is: {similarity_score}")


if __name__ == "__main__":
    main()
