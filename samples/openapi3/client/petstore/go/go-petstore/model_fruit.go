/*
 * fruity
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore
import (
    "encoding/json"
    "errors"
)

// Fruit struct for Fruit
type Fruit struct {
    Apple *Apple
    Banana *Banana
}

func (dst *Fruit) UnmarshalJSON(data []byte) error {
    var err error
    err = json.Unmarshal(data, dst.Apple);
    if err == nil {
        return nil
    } else {
        dst.Apple = nil
    }
    err = json.Unmarshal(data, dst.Banana);
    if err == nil {
        return nil
    } else {
        dst.Banana = nil
    }
    return err
}

func (src *Fruit) MarshalJSON() ([]byte, error) {
    if src.Apple != nil {
        return json.Marshal(src.Apple)
    }
    if src.Banana != nil {
        return json.Marshal(src.Banana)
    }
    return nil, errors.New("No field set in Fruit")
}

