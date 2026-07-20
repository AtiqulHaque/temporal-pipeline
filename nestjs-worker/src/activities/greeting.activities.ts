export function getGreetingActivity(name: string): string {
  return `Hello, ${name} — processed by NestJS worker at ${new Date().toISOString()}`;
}
