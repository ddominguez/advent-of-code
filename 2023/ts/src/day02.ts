async function main() {
	const input = Bun.file("../input/02.txt");
	const data = await input.text();
	if (Bun.argv[2] === "part2") {
		console.log(part2(data.trim()));
	} else {
		console.log(part1(data.trim()));
	}
}

export function part1(data: string) {
	let result = 0;
	const maxColorCounts = new Map([
		["red", 12],
		["green", 13],
		["blue", 14],
	]);

	for (const line of data.split("\n")) {
		const [game, sets] = line.split(": ");
		let isPossible = true;

		for (const set of sets.split("; ")) {
			if (!isPossible) {
				break;
			}
			for (const cube of set.split(", ")) {
				const [count, color] = cube.split(" ");
				const countInt = Number.parseInt(count);
				const maxCount = maxColorCounts.get(color);
				if (maxCount && maxCount < countInt) {
					isPossible = false;
					break;
				}
			}
		}
		if (isPossible) {
			result += Number.parseInt(game.replace(/\D/g, ""));
		}
	}

	return result;
}

export function part2(data: string) {
	let result = 0;

	for (const line of data.split("\n")) {
		const [_, sets] = line.split(": ");
		const minColorCount = new Map([
			["red", 0],
			["green", 0],
			["blue", 0],
		]);

		for (const set of sets.split("; ")) {
			for (const cube of set.split(", ")) {
				const [count, color] = cube.split(" ");
				const countInt = Number.parseInt(count);
				const minCount = minColorCount.get(color);
				if (minCount !== undefined) {
					minColorCount.set(color, Math.max(minCount, countInt));
				}
			}
		}
		result += Array.from(minColorCount.values()).reduce(
			(acc, curr) => acc * curr,
			1,
		);
	}

	return result;
}

if (import.meta.path === Bun.main) {
	main();
}
