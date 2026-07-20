export function logResult(greeting: string): string {
  return `Log result message: ${greeting} — processed by NestJS worker at ${new Date().toISOString()}`;
}
