/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow strict-local
 */

import React from 'react';
import type { Node } from 'react';
import {
  SafeAreaView,
  ScrollView,
  StatusBar,
  StyleSheet,
  Text,
  useColorScheme,
  View,
  Button,
  Modal
} from 'react-native';

import {
  Colors,
} from 'react-native/Libraries/NewAppScreen';

import FoodCalendar from './containers/food_calendar';
import { FoodForm, RecordForm } from './containers/forms'

const backendUrl = 'http://192.168.1.101:8080'

/* $FlowFixMe[missing-local-annot] The type annotation(s) required by Flow's
 * LTI update could not be added via codemod */
const Section = ({ children, title }): Node => {
  const isDarkMode = useColorScheme() === 'dark';
  return (
    <View style={styles.sectionContainer}>
      <Text
        style={[
          styles.sectionTitle,
          {
            color: isDarkMode ? Colors.white : Colors.black,
          },
        ]}>
        {title}
      </Text>
      <Text
        style={[
          styles.sectionDescription,
          {
            color: isDarkMode ? Colors.light : Colors.dark,
          },
        ]}>
        {children}
      </Text>
    </View>
  );
};

const App: () => Node = () => {

  const defaultRecord = {
    Name: "",
    EatingDate: "",
    EatenQuantity: 0,
    SatisfactionScore: 0,
    isModifying: false
  };
  const [isCreatingFood, onChangeIsCreatingFood] = React.useState(false);
  const [curRecord, onChangeCurRecord] = React.useState(defaultRecord);

  const isDarkMode = useColorScheme() === 'dark';
  const backgroundStyle = {
    backgroundColor: isDarkMode ? Colors.darker : Colors.lighter,
  };

  return (
    <SafeAreaView style={backgroundStyle}>
      <StatusBar
        barStyle={isDarkMode ? 'light-content' : 'dark-content'}
        backgroundColor={backgroundStyle.backgroundColor}
      />
      <ScrollView
        contentInsetAdjustmentBehavior="automatic"
        style={backgroundStyle}>
        {/* <Header /> */}
        <View
          style={{
            backgroundColor: isDarkMode ? Colors.black : Colors.white,
          }}
        >
          <Button
            onPress={() => { onChangeIsCreatingFood(true) }}
            title="Add Food"
            color="#841584"
            accessibilityLabel="This button is used to add food"
          />
          <Modal visible={isCreatingFood}>
            <FoodForm closeHandle={onChangeIsCreatingFood} backendUrl={backendUrl} />
          </Modal>

          <Section title="Food Calendar">
            <FoodCalendar
              createForm={(rec) => { onChangeCurRecord(rec) }}
              backendUrl={backendUrl}
            />
            <Modal visible={curRecord.EatingDate !== ""}>
              <RecordForm
                record={curRecord}
                closeHandle={() => { onChangeCurRecord(defaultRecord) }}
                backendUrl={backendUrl}
              />
            </Modal>
          </Section>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  sectionContainer: {
    marginTop: 32,
    paddingHorizontal: 24,
  },
  sectionTitle: {
    fontSize: 24,
    fontWeight: '600',
  },
  sectionDescription: {
    marginTop: 8,
    fontSize: 18,
    fontWeight: '400',
  },
  highlight: {
    fontWeight: '700',
  },
});

export default App;
