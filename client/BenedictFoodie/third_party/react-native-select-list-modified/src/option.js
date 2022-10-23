import React from 'react';
import { StyleSheet, View, Text } from 'react-native';


const Option = ({ children }) => {
  return (
    <View style={[styles.option]}>
      <Text style={[styles.text]}>{children}</Text>
    </View>
  )
}

const styles = StyleSheet.create({
  option: {
    justifyContent: 'center',
    borderBottomWidth: 1,
    borderBottomColor: '#cccccc',
    marginHorizontal: 10,
    paddingVertical: 10,
  },
  text: {
    paddingHorizontal: 5,
  },
});

module.exports = Option