export function searchUserMessage(greeting: string): string {
  return `User search message: ${greeting} — processed by NestJS worker at ${new Date().toISOString()}`;
}
