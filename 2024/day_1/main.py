def get_locations() -> tuple[list[int], list[int]]:
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


def first_challenge(first_location_list: list[int], second_location_list: list[int]) -> int:
    location_distance = 0
    for first_location, second_location in zip(first_location_list, second_location_list):
        if first_location > second_location:
            location_distance += first_location - second_location
        elif second_location > first_location:
            location_distance += second_location - first_location

    return location_distance # 1223326


def main():
    first_location_list, second_location_list = get_locations()
    first_challenge_result = first_challenge(first_location_list,second_location_list)
    print(first_challenge_result)


if __name__ == "__main__":
    main()
