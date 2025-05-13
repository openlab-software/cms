import { HelloWorld } from "@cms/ts-common";

export default async function Page() {
  const response = await fetch("http://localhost:8080/v1");
  const helloWorld: HelloWorld = await response.json();

  return <div>{helloWorld.message}</div>;
}
