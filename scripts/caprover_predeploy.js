var preDeployFunction = async function (captainAppObj, dockerUpdateObject) {
  const DockerApi = require("./built/docker/DockerApi");
  const api = new DockerApi.default();
  const run = async (args) => {
    const imageName = dockerUpdateObject.TaskTemplate.ContainerSpec.Image;
    const env = captainAppObj.envVars.map((kv) => kv.key + "=" + kv.value);
    const config = {
      Env: env,
      HostConfig: { AutoRemove: true, NetworkMode: captainAppObj.networks[0] },
    };

    await api.dockerode.run(imageName, args, process.stdout, config);
  };

  try {
    // nuke, reseed
    // to check output,  `docker service logs captain-captain --since 10m --follow`
    const env = {};
    captainAppObj.envVars.map((kv) => (env[kv.key] = kv.value));

    // auto use caprover details
    await run([
      "/app/bin/migrate",
      "-database",
      `postgres://${env.RENTHOME_DATABASE_USER}:${env.RENTHOME_DATABASE_PASS}@${env.RENTHOME_DATABASE_HOST}:${env.RENTHOME_DATABASE_PORT}/${env.RENTHOME_DATABASE_NAME}?sslmode=disable`,
      "-path",
      "/app/migrations",
      "drop",
      "-f",
    ]);
    await run([
      "/app/bin/migrate",
      "-database",
      `postgres://${env.RENTHOME_DATABASE_USER}:${env.RENTHOME_DATABASE_PASS}@${env.RENTHOME_DATABASE_HOST}:${env.RENTHOME_DATABASE_PORT}/${env.RENTHOME_DATABASE_NAME}?sslmode=disable`,
      "-path",
      "/app/migrations",
      "up",
    ]);
    await run(["/app/main", "db", "--seed"]);

    // migrate only
    // await run([
    // 	"/app/bin/migrate",
    // 	"-database",
    // 	`postgres://${env.RENTHOME_DATABASE_USER}:${env.RENTHOME_DATABASE_PASS}@${env.RENTHOME_DATABASE_HOST}:${env.RENTHOME_DATABASE_PORT}/${env.RENTHOME_DATABASE_NAME}?sslmode=disable`,
    // 	"-path",
    // 	"/app/migrations",
    // 	"up",
    // ])
  } catch (err) {
    console.error(err);
  }
  return dockerUpdateObject;
};
