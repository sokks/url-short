import productionConfig from './configureStore.prod';
import developmentConfig from './configureStore.dev';

let config = null;

if (process.env.NODE_ENV === 'production') {
  config = productionConfig;
} else {
  config = developmentConfig;
}

export default config;
