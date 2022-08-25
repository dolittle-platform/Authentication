import RubikRegularWoff from './Rubik-Regular.woff2';
import RubikBoldWoff from './Rubik-Bold.woff2';
import RubikMediumWoff from './Rubik-Medium.woff2';
import RubikLightWoff from './Rubik-Light.woff2';

const RubikRegular = {
    fontFamily: 'Rubik',
    src: `local('Rubik'), local('Rubik-Regular'), url(${RubikRegularWoff}) format('woff2')`,
    fontDisplay: 'swap',
};

const RubikBold = {
    fontFamily: 'Rubik-Bold',
    src: `local('Rubik-Bold'), url(${RubikBoldWoff}) format('woff2')`,
    fontDisplay: 'swap',
};

const RubikMedium = {
    fontFamily: 'Rubik-Medium',
    src: `local('Rubik-Medium'), url(${RubikMediumWoff}) format('woff2')`,
    fontDisplay: 'swap',
};

const RubikLight = {
    fontFamily: 'Rubik-Light',
    src: `local('Rubik-Light'), url(${RubikLightWoff}) format('woff2')`,
    fontDisplay: 'swap',
};

export const Rubik = {
    Regular: RubikRegular,
    Bold: RubikBold,
    Medium: RubikMedium,
    Light: RubikLight,
};
