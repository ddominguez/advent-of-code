<?php
function get_puzzle_contents(int $day): string {
    $d = str_pad($day, 2, '0', STR_PAD_LEFT);
    $contents = trim(file_get_contents("../input/$d.txt"));
    return $contents;
}

function main(array $context): void {
    $contents =  get_puzzle_contents($context['day']);
    switch($context['cmd']) {
        case 'part1':
            echo $context['part1']($contents) . PHP_EOL;
            break;
        case 'part2':
            echo $context['part2']($contents) . PHP_EOL;
            break;
        case 'test':
            $context['tests']();
            break;
        default:
            echo "Unknown command: " . $context['cmd'] . PHP_EOL;
            break;
    }
}

function expect(mixed $actual, mixed $expected, string $func_name): void {
    echo "testing $func_name";
    if ($actual === $expected) {
        echo "...OK" . PHP_EOL;
    } else {
        echo "...FAIL - expected $expected, got $actual" . PHP_EOL;
    }
}
