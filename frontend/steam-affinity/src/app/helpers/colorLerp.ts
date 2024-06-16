export interface Color {
	r: number;
	g: number;
	b: number;
}

function lerp(start: number, end: number, t: number): number {
	return start + (end - start) * t;
}

function colorLerp(startColor: Color, endColor: Color, t: number): Color {
	let r = Math.round(lerp(startColor.r, endColor.r, t))
	let g = Math.round(lerp(startColor.g, endColor.g, t))
	let b = Math.round(lerp(startColor.b, endColor.b, t))

	return { r, g, b };
}

export function getColorForValue(value: number, min: number, max: number, startColor: Color, endColor: Color): Color {
	if (value < min) {
		value = min;
	}
	if (value > max) {
		value = max;
	}

	let t = (value - min) / (max - min);
	return colorLerp(startColor, endColor, t);
}

export function colorToStyle(color: Color): string {
	return `rgb(${color.r}, ${color.g}, ${color.b})`
}
