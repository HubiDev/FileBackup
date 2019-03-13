using System.ComponentModel;
using System.Runtime.CompilerServices;

namespace FileBackup.Framework
{
    class ViewModelBase : INotifyPropertyChanged
    {
        /// <summary>
        /// Event Handler for the property changed event
        /// </summary>
        public event PropertyChangedEventHandler PropertyChanged;

        /// <summary>
        /// Invokes the property changed event
        /// </summary>
        /// <param name="propertyName">The of the property for which the event should be invoked</param>
        protected void OnPropertyChanged([CallerMemberName] string propertyName = null)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }
    }
}
