import convict from 'convict';
import dotenv from 'dotenv';

/** The manifest describing the config */
const configManifest = {
  http: {
    port: {
      default: 3000,
      doc: 'The port to listen on',
      env: 'PORT',
      format: 'port'
    }
  }
};

/**
 * Load and return the configuration for the app
 */
export function loadConfig(): convict.Config<any> {
  dotenv.config();
  const config = convict(configManifest);
  return config;
}
